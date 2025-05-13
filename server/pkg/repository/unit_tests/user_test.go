package unit_tests

import (
	"github.com/vujanic79/golang-react-todo-app/pkg/repository"
	"testing"
)

func TestMapDbUserToUser(t *testing.T) {
	dbU := generateDbUser("John", "Doe", "john.doe@gmail.com")
	want := generateUser(dbU)

	u := repository.MapDbUserToUser(dbU)
	areEqual := checkUserEquality(want, u)

	if !areEqual {
		t.Errorf("MapDbUserToUser(dbU) = %v, want %v", u, want)
	}
}
