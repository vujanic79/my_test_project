package unit_tests

import (
	"github.com/vujanic79/golang-react-todo-app/pkg/repository"
	"testing"
)

func TestMapDbTaskToTask(t *testing.T) {
	dbT := generateDbTask("Task 1", "Task 1 description", "ACTIVE")
	want := generateTask(dbT)

	task := repository.MapDbTaskToTask(dbT)
	areEqual := checkTaskEquality(want, task)

	if !areEqual {
		t.Errorf("MapDbTaskToTask(dbT) = %v, want %v", task, want)
	}
}

func TestMapDbTasksToTasks(t *testing.T) {
	dbTs := generateDbTasks()
	want := generateTasks(dbTs)

	ts := repository.MapDbTasksToTasks(dbTs)
	areEqual := checkTasksEquality(want, ts)

	if !areEqual {
		t.Errorf("MapDbTasksToTasks(dbTs) = %v, want %v", ts, want)
	}
}
