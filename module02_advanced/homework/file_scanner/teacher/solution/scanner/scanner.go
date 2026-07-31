package scanner

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var (
	ErrInvalidRoot        = errors.New("root must be an existing directory")
	ErrEmptyNeedle        = errors.New("search text must not be empty")
	ErrInvalidWorkerCount = errors.New("worker count must be greater than zero")
)

type Result struct {
	Files         int
	MatchingLines int
}

type fileResult struct {
	matchingLines int
	err           error
}

func Scan(parent context.Context, root, needle string, workers int) (Result, error) {
	if needle == "" {
		return Result{}, ErrEmptyNeedle
	}
	if workers < 1 {
		return Result{}, ErrInvalidWorkerCount
	}
	info, err := os.Stat(root)
	if err != nil || !info.IsDir() {
		return Result{}, fmt.Errorf("%w: %s", ErrInvalidRoot, root)
	}
	if err := parent.Err(); err != nil {
		return Result{}, err
	}

	ctx, cancel := context.WithCancel(parent)
	defer cancel()
	jobs := make(chan string)
	results := make(chan fileResult)
	walkDone := make(chan error, 1)

	go func() {
		defer close(jobs)
		walkDone <- filepath.WalkDir(root, func(path string, entry fs.DirEntry, walkErr error) error {
			if walkErr != nil {
				return walkErr
			}
			if !entry.Type().IsRegular() {
				return nil
			}
			select {
			case <-ctx.Done():
				return ctx.Err()
			case jobs <- path:
				return nil
			}
		})
	}()

	var workerGroup sync.WaitGroup
	for i := 0; i < workers; i++ {
		workerGroup.Add(1)
		go func() {
			defer workerGroup.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case filename, ok := <-jobs:
					if !ok {
						return
					}
					matches, scanErr := countMatchingLines(filename, needle)
					select {
					case <-ctx.Done():
						return
					case results <- fileResult{matchingLines: matches, err: scanErr}:
					}
					if scanErr != nil {
						return
					}
				}
			}
		}()
	}

	go func() {
		workerGroup.Wait()
		close(results)
	}()

	var result Result
	var firstErr error
	for file := range results {
		if file.err != nil && firstErr == nil {
			firstErr = file.err
			cancel()
			continue
		}
		result.Files++
		result.MatchingLines += file.matchingLines
	}
	walkErr := <-walkDone
	if firstErr != nil {
		return result, firstErr
	}
	if err := parent.Err(); err != nil {
		return result, err
	}
	if walkErr != nil {
		return result, fmt.Errorf("walk %s: %w", root, walkErr)
	}
	return result, nil
}

func countMatchingLines(filename, needle string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, fmt.Errorf("open %s: %w", filename, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, 64*1024), 1024*1024)
	matches := 0
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), needle) {
			matches++
		}
	}
	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("scan %s: %w", filename, err)
	}
	return matches, nil
}
