package tasks

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestTask struct {
	err error
}

func (tt *TestTask) Run() error {
	return tt.err
}

func TestProcess1(t *testing.T) {

	testTask1 := &TestTask{err: errors.New("error")}
	testTask2 := &TestTask{err: nil}
	testTask3 := &TestTask{err: errors.New("other error")}

	testTask := []Task{testTask1, testTask2, testTask3}

	result := Process(testTask)
	assert.Len(t, result, 2)
}

func TestProcess2(t *testing.T) {

	testTask1 := &TestTask{err: errors.New("error")}
	testTask2 := &TestTask{err: nil}
	testTask3 := &TestTask{err: errors.New("error")}

	testTask := []Task{testTask1, testTask2, testTask3}

	result := Process(testTask)
	assert.Len(t, result, 2)
}

func TestProcess3(t *testing.T) {

	testTask1 := &TestTask{err: errors.New("error")}
	testTask2 := &TestTask{err: errors.New("error")}
	testTask3 := &TestTask{err: errors.New("error")}

	testTask := []Task{testTask1, testTask2, testTask3}

	result := Process(testTask)
	assert.Len(t, result, 3)
}
