package unit_tests

import (
	"github.com/google/uuid"
	"github.com/vujanic79/golang-react-todo-app/pkg/domain"
	"github.com/vujanic79/golang-react-todo-app/pkg/internal/database"
	"slices"
	"time"
)

func generateDbUser(firstName string, lastName string, email string) (dbU database.AppUser) {
	return database.AppUser{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
	}
}

func generateUser(dbU database.AppUser) (u domain.User) {
	return domain.User{
		ID:        dbU.ID,
		CreatedAt: dbU.CreatedAt,
		UpdatedAt: dbU.UpdatedAt,
		FirstName: dbU.FirstName,
		LastName:  dbU.LastName,
		Email:     dbU.Email,
	}
}

func checkUserEquality(want domain.User, u domain.User) bool {
	return want.ID == u.ID &&
		want.CreatedAt == u.CreatedAt &&
		want.UpdatedAt == u.UpdatedAt &&
		want.FirstName == u.FirstName &&
		want.LastName == u.LastName &&
		want.Email == u.Email
}

func generateDbTask(title string, description string, status string) (dbT database.Task) {
	return database.Task{
		ID:               uuid.New(),
		CreatedAt:        time.Now().UTC(),
		UpdatedAt:        time.Now().UTC(),
		Title:            title,
		Description:      description,
		Status:           status,
		CompleteDeadline: time.Now().UTC().Add(1 * time.Hour),
		UserID:           uuid.New(),
	}
}

func generateTask(dbT database.Task) (t domain.Task) {
	return domain.Task{
		ID:               dbT.ID,
		CreatedAt:        dbT.CreatedAt,
		UpdatedAt:        dbT.UpdatedAt,
		Title:            dbT.Title,
		Description:      dbT.Description,
		Status:           dbT.Status,
		CompleteDeadline: dbT.CompleteDeadline,
		UserID:           dbT.UserID,
	}
}

func generateDbTasks() (dbTs []database.Task) {
	return []database.Task{
		generateDbTask("Task1", "Description1", "ACTIVE"),
		generateDbTask("Task2", "Description2", "PENDING"),
	}
}

func generateTasks(dbTs []database.Task) (ts []domain.Task) {
	tasks := make([]domain.Task, len(dbTs))
	for i, dbTask := range dbTs {
		tasks[i] = generateTask(dbTask)
	}
	return tasks
}

func checkTaskEquality(want domain.Task, t domain.Task) bool {
	return want.ID == want.ID &&
		want.CreatedAt == t.CreatedAt &&
		want.UpdatedAt == t.UpdatedAt &&
		want.Title == t.Title &&
		want.Description == t.Description &&
		want.Status == t.Status &&
		want.CompleteDeadline == t.CompleteDeadline &&
		want.UserID == want.UserID
}

func checkTasksEquality(want []domain.Task, ts []domain.Task) bool {
	return slices.Equal(want, ts)
}

func generateTaskStatuses(dbtss []string) (tss []domain.TaskStatus) {
	tss = make([]domain.TaskStatus, len(dbtss))
	for i, dbTs := range dbtss {
		tss[i] = domain.TaskStatus{Status: dbTs}
	}
	return tss
}

func checkTaskStatusEquality(want domain.TaskStatus, ts domain.TaskStatus) bool {
	return want.Status == ts.Status
}

func checkTaskStatusesEquality(want []domain.TaskStatus, tss []domain.TaskStatus) bool {
	return slices.Equal(want, tss)
}
