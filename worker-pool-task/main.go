package main

import (
	"context"
	"errors"
	"math/rand"
	"sync"
)

var taskRunnerError = errors.New("Task runner error")

type Task struct {
}

type TaskRunner interface {
	run() (error, bool)
}

func (t *Task) run() error {

	if rand.Intn(2) == 0 {
		return taskRunnerError
	} else {
		return nil
	}

}

func start(n int, m int, tasks []Task) {
	var wg sync.WaitGroup

	errorChannel := make(chan error)
	taskChannel := make(chan Task)
	ctx, cancel := context.WithCancel(context.Background())

	for i := 0; i < n; i++ {
		wg.Add(1)
		go handleGoroutin(ctx, taskChannel, errorChannel, &wg)
	}

	go func() {
		for i := 0; i < len(tasks); i++ {
			select {
			case <-ctx.Done():
				return
			case taskChannel <- tasks[i]:
			}
		}
	}()

	go func() {
		var errorsCounter int = 0
		for range errorChannel {
			errorsCounter++
			if m == errorsCounter {
				cancel()
			}
		}
	}()

	wg.Wait()

	close(errorChannel)
}

func handleGoroutin(ctx context.Context, taskChannel chan Task, errorChannel chan error, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		case task, ok := <-taskChannel:
			if !ok {
				close(taskChannel)
				return
			}
			err := task.run()

			if err != nil {
				errorChannel <- err
			}
		}
	}
}

func main() {

	var countTask int = 30
	var errorLimit int = 5

	tasks := make([]Task, countTask, 30)

	start(countTask, errorLimit, tasks)
}
