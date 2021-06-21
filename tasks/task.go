package tasks

import (
	"sync"
)

type Task interface {
	Run() error
}

//Task1 -> err
//Task2 -> nil
//Task3 -> err
//
//process([Task1, Task2, Task3]) -> [err1, err2]

func Process(tasks []Task) []error {
	var errors []error
	var wg sync.WaitGroup
	errBuf := make(chan error, len(tasks))

	for _, task := range tasks {
		wg.Add(1)
		go func(task Task) {
			defer wg.Done()
			err := task.Run()
			errBuf <- err
		}(task)

	}

	wg.Wait()
	close(errBuf)

	for err := range errBuf {
		if err != nil {
			errors = append(errors, err)

		}
	}

	return errors
}
