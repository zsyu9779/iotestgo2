package server

import (
	"context"
	"errors"
	"io"
	"testing"
	"time"

	pb "iotestgo/module05_grpc/project_distributed_compute/proto/computepb"

	"google.golang.org/grpc/metadata"
)

type fakeProcessStream struct {
	pb.DistributedCompute_ProcessServer
	ctx     context.Context
	inputs  []*pb.ComputeTask
	outputs []*pb.ComputeResult
	index   int
	sendErr error
}

func (f *fakeProcessStream) Context() context.Context {
	if f.ctx == nil {
		return context.Background()
	}
	return f.ctx
}

func (f *fakeProcessStream) Recv() (*pb.ComputeTask, error) {
	if f.index >= len(f.inputs) {
		return nil, io.EOF
	}
	task := f.inputs[f.index]
	f.index++
	return task, nil
}

func (f *fakeProcessStream) Send(result *pb.ComputeResult) error {
	if f.sendErr != nil {
		return f.sendErr
	}
	f.outputs = append(f.outputs, result)
	return nil
}

func (f *fakeProcessStream) SendHeader(metadata.MD) error { return nil }
func (f *fakeProcessStream) SetHeader(metadata.MD) error  { return nil }
func (f *fakeProcessStream) SetTrailer(metadata.MD)       {}

func TestServiceProcess(t *testing.T) {
	stream := &fakeProcessStream{
		inputs: []*pb.ComputeTask{
			{TaskId: "sum", Operation: "sum", Numbers: []int64{1, 2, 3}},
			{TaskId: "bad", Operation: "p99", Numbers: []int64{1, 2, 3}},
		},
	}

	err := NewService(2).Process(stream)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(stream.outputs) != 2 {
		t.Fatalf("expected 2 outputs, got %d", len(stream.outputs))
	}

	statuses := map[string]string{}
	for _, result := range stream.outputs {
		statuses[result.GetTaskId()] = result.GetStatus()
	}
	if statuses["sum"] != "done" {
		t.Fatalf("expected sum done, got %q", statuses["sum"])
	}
	if statuses["bad"] != "error" {
		t.Fatalf("expected bad error, got %q", statuses["bad"])
	}
}

func TestServiceProcessReturnsSendError(t *testing.T) {
	wantErr := errors.New("send failed")
	inputs := make([]*pb.ComputeTask, 200)
	for i := range inputs {
		inputs[i] = &pb.ComputeTask{TaskId: "sum", Operation: "sum", Numbers: []int64{1, 2, 3}}
	}
	stream := &fakeProcessStream{
		inputs:  inputs,
		sendErr: wantErr,
	}

	done := make(chan error, 1)
	go func() {
		done <- NewService(2).Process(stream)
	}()

	select {
	case err := <-done:
		if !errors.Is(err, wantErr) {
			t.Fatalf("expected send error %v, got %v", wantErr, err)
		}
	case <-time.After(time.Second):
		t.Fatal("Process did not return after stream.Send failed")
	}
}
