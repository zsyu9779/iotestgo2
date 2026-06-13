package server

import (
	"io"
	"log"
	"sync"

	"iotestgo/module05_grpc/project_distributed_compute/internal/engine"
	pb "iotestgo/module05_grpc/project_distributed_compute/proto/computepb"
)

type Service struct {
	pb.UnimplementedDistributedComputeServer
	WorkerCount int
}

func NewService(workerCount int) *Service {
	if workerCount <= 0 {
		workerCount = 1
	}
	return &Service{WorkerCount: workerCount}
}

func (s *Service) Process(stream pb.DistributedCompute_ProcessServer) error {
	tasksCh := make(chan *pb.ComputeTask, 100)
	resultsCh := make(chan *pb.ComputeResult, 100)
	stopCh := make(chan struct{})
	var stopOnce sync.Once
	stop := func() {
		stopOnce.Do(func() {
			close(stopCh)
		})
	}

	var workers sync.WaitGroup
	for i := 0; i < s.WorkerCount; i++ {
		workers.Add(1)
		go func(workerID int) {
			defer workers.Done()
			for {
				select {
				case <-stopCh:
					return
				case task, ok := <-tasksCh:
					if !ok {
						return
					}
					result := computeTask(task)
					select {
					case resultsCh <- result:
						log.Printf("[Worker-%d] task=%s op=%s", workerID, task.GetTaskId(), task.GetOperation())
					case <-stopCh:
						return
					}
				}
			}
		}(i + 1)
	}

	sendDone := make(chan error, 1)
	go func() {
		for result := range resultsCh {
			if err := stream.Send(result); err != nil {
				stop()
				sendDone <- err
				return
			}
		}
		sendDone <- nil
	}()

	for {
		task, err := stream.Recv()
		if err == io.EOF {
			close(tasksCh)
			workers.Wait()
			close(resultsCh)
			return <-sendDone
		}
		if err != nil {
			close(tasksCh)
			workers.Wait()
			close(resultsCh)
			<-sendDone
			return err
		}
		select {
		case tasksCh <- task:
		case <-stopCh:
			close(tasksCh)
			workers.Wait()
			close(resultsCh)
			return <-sendDone
		}
	}
}

func computeTask(task *pb.ComputeTask) *pb.ComputeResult {
	value, err := engine.Compute(engine.ParseOperation(task.GetOperation()), task.GetNumbers())
	result := &pb.ComputeResult{
		TaskId:    task.GetTaskId(),
		Operation: task.GetOperation(),
	}
	if err != nil {
		result.Status = "error"
		result.Message = err.Error()
		return result
	}
	result.Status = "done"
	result.Value = value
	return result
}
