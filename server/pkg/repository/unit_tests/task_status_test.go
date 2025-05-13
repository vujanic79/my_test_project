package unit_tests

import (
	"github.com/vujanic79/golang-react-todo-app/pkg/domain"
	"github.com/vujanic79/golang-react-todo-app/pkg/repository"
	"testing"
)

func TestMapDbTaskStatusToTaskStatus(t *testing.T) {
	status := "PENDING"
	want := domain.TaskStatus{Status: status}

	ts := repository.MapDbTaskStatusToTaskStatus(status)
	areEqual := checkTaskStatusEquality(want, ts)

	if !areEqual {
		t.Errorf("MapDbTaskStatusToTaskStatus(status) = %v, want %v", ts.Status, want.Status)
	}
}

func TestMapDbTaskStatusesToTaskStatuses(t *testing.T) {
	statuses := []string{"PENDING", "ACTIVE", "COMPLETED"}
	want := generateTaskStatuses(statuses)

	tss := repository.MapDbTaskStatusesToTaskStatuses(statuses)
	areEqual := checkTaskStatusesEquality(want, tss)

	if !areEqual {
		t.Errorf("MapDbTaskStatusesToTaskStatuses(statuses) = %v, want %v", tss, want)
	}
}
