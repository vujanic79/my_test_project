package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/vujanic79/golang-react-todo-app/pkg/domain"
	"io"
	"net/http"
	"os"
	"testing"
	"time"
)

func TestGetTasksByUserId(t *testing.T) {
	p := os.Getenv("PORT")

	userId, err := uuid.Parse("32c5982d-b0a7-4756-94f0-11a468ffe05d")
	if err != nil {
		t.Fatalf("Parsing user id error: %v", err)
	}
	body := domain.GetTasksByUserIdParams{
		UserID: userId,
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		t.Fatalf("Could not marshal body: %v", err)
	}

	req, err := http.NewRequest("POST",
		fmt.Sprintf("http://127.0.0.1:%s/todo/tasks-by-user", p), bytes.NewBuffer(jsonBody))

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		t.Fatalf("Executing POST request error: %v", err)
	}

	defer func(body io.ReadCloser) {
		err := body.Close()
		if err != nil {
			t.Errorf("Closing response body error: %v", err)
		}
	}(res.Body)

	if http.StatusOK != res.StatusCode {
		t.Errorf("Expected status: %d, actual: %d", http.StatusOK, res.StatusCode)
	}

	decoder := json.NewDecoder(res.Body)
	var actual []domain.Task
	err = decoder.Decode(&actual)
	if err != nil {
		t.Fatalf("Decoding response body error: %v", err)
	}

	id, err := uuid.Parse("012d38ae-f234-46f8-bea8-4a3f43b0abb0")
	if err != nil {
		t.Fatalf("Parsing task id error: %v", err)
	}
	var expected = []domain.Task{
		{
			ID:               id,
			CreatedAt:        time.Now().UTC(),
			UpdatedAt:        time.Now().UTC(),
			Title:            "Task 001",
			Description:      "Task description 001",
			Status:           "ACTIVE",
			CompleteDeadline: time.Now().UTC().Add(7 * 24 * time.Hour),
			UserID:           userId,
		},
	}

	if !areTaskSlicesEqual(expected, actual) {
		t.Errorf("Response body content error, expected: %s, actual: %s", expected, actual)
	}
}

// Trying to get tasks for non-existing userId: 32c5982d-b0a7-4756-94f0-11a468ffe05c
func TestGetTasksByUserIdEmptyResponse(t *testing.T) {
	p := os.Getenv("PORT")

	userId, err := uuid.Parse("32c5982d-b0a7-4756-94f0-11a468ffe05c")
	if err != nil {
		t.Fatalf("Parsing user id error: %v", err)
	}
	body := domain.GetTasksByUserIdParams{
		UserID: userId,
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		t.Fatalf("Could not marshal body: %v", err)
	}

	req, err := http.NewRequest("POST",
		fmt.Sprintf("http://127.0.0.1:%s/todo/tasks-by-user", p), bytes.NewBuffer(jsonBody))

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		t.Fatalf("Executing POST request error: %v", err)
	}

	defer func(body io.ReadCloser) {
		err := body.Close()
		if err != nil {
			t.Errorf("Closing response body error: %v", err)
		}
	}(res.Body)

	if http.StatusOK != res.StatusCode {
		t.Errorf("Expected status: %d, actual: %d", http.StatusOK, res.StatusCode)
	}

	decoder := json.NewDecoder(res.Body)
	var actual []domain.Task
	err = decoder.Decode(&actual)
	if err != nil {
		t.Fatalf("Decoding response body error: %v", err)
	}

	var expected []domain.Task

	if !areTaskSlicesEqual(expected, actual) {
		t.Errorf("Response body content error, expected: %s, actual: %s", expected, actual)
	}
}

func TestCreateTaskPass(t *testing.T) {
	p := os.Getenv("PORT")

	completeDeadline := time.Now().UTC().Add(7 * 24 * time.Hour).Format("2006-01-02T15:04:05.999999Z")
	body := domain.CreateTaskParams{
		Title:            "Do some exercises!",
		Description:      "It's really good for your wellbeing",
		Status:           "PAUSED",
		CompleteDeadline: completeDeadline,
		UserEmail:        "milanmvujanic@gmail.com",
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		t.Fatalf("Could not marshal body: %v", err)
	}

	req, err := http.NewRequest("POST",
		fmt.Sprintf("http://127.0.0.1:%s/todo/tasks", p),
		bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatalf("Creating POST request error: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		t.Fatalf("Executing POST request error: %v", err)
	}

	defer func(body io.ReadCloser) {
		err := body.Close()
		if err != nil {
			t.Errorf("Closing response body error: %v", err)
		}
	}(res.Body)

	if http.StatusCreated != res.StatusCode {
		t.Errorf("Expected status: %d, actual: %d", http.StatusCreated, res.StatusCode)
	}

	decoder := json.NewDecoder(res.Body)
	var actual domain.Task
	err = decoder.Decode(&actual)
	if err != nil {
		t.Fatalf("Decoding response body error: %v", err)
	}

	userId, err := uuid.Parse("32c5982d-b0a7-4756-94f0-11a468ffe05d")
	if err != nil {
		t.Fatalf("Parsing user id error: %v", err)
	}
	var expected = domain.Task{
		ID:               actual.ID,
		CreatedAt:        time.Now().UTC(),
		UpdatedAt:        time.Now().UTC(),
		Title:            "Do some exercises!",
		Description:      "It's really good for your wellbeing",
		Status:           "PAUSED",
		CompleteDeadline: time.Now().UTC().Add(7 * 24 * time.Hour),
		UserID:           userId,
	}

	if !areTasksEqual(expected, actual) {
		t.Errorf("Response body content error, expected: %s, actual: %s", expected, actual)
	}
}

// Trying to create task for non-existing user email: somebody@email.com
func TestCreateTaskFail(t *testing.T) {
	p := os.Getenv("PORT")

	completeDeadline := time.Now().UTC().Add(7 * 24 * time.Hour).Format("2006-01-02T15:04:05.999999Z")
	body := domain.CreateTaskParams{
		Title:            "Do some exercises!",
		Description:      "It's really good for your wellbeing",
		Status:           "PAUSED",
		CompleteDeadline: completeDeadline,
		UserEmail:        "somebody@email.com",
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		t.Fatalf("Could not marshal body: %v", err)
	}

	req, err := http.NewRequest("POST",
		fmt.Sprintf("http://127.0.0.1:%s/todo/tasks", p),
		bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatalf("Creating POST request error: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		t.Fatalf("Executing POST request error: %v", err)
	}

	defer func(body io.ReadCloser) {
		err := body.Close()
		if err != nil {
			t.Errorf("Closing response body error: %v", err)
		}
	}(res.Body)

	if http.StatusInternalServerError != res.StatusCode {
		t.Errorf("Expected status: %d, actual: %d", http.StatusInternalServerError, res.StatusCode)
	}
}

func TestUpdateTaskPass(t *testing.T) {
	p := os.Getenv("PORT")

	idStr := "012d38ae-f234-46f8-bea8-4a3f43b0abb0"
	completeDeadline := time.Now().UTC().Add(7 * 24 * time.Hour).Format("2006-01-02T15:04:05.999999Z")
	body := domain.UpdateTaskParams{
		Title:            "Task 001 updated",
		Description:      "Task 001 description updated",
		Status:           "PAUSED",
		CompleteDeadline: completeDeadline,
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		t.Fatalf("Could not marshal body: %v", err)
	}

	req, err := http.NewRequest("PUT",
		fmt.Sprintf("http://127.0.0.1:%s/todo/tasks/%s", p, idStr),
		bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatalf("Creating PUT request error: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		t.Fatalf("Executing PUT request error: %v", err)
	}

	defer func(body io.ReadCloser) {
		err := body.Close()
		if err != nil {
			t.Errorf("Closing response body error: %v", err)
		}
	}(res.Body)

	if http.StatusOK != res.StatusCode {
		t.Errorf("Expected status: %d, actual: %d", http.StatusOK, res.StatusCode)
	}

	decoder := json.NewDecoder(res.Body)
	var actual domain.Task
	err = decoder.Decode(&actual)
	if err != nil {
		t.Fatalf("Decoding response body error: %v", err)
	}

	var expected = domain.Task{
		ID:               actual.ID,
		CreatedAt:        time.Now().UTC(),
		UpdatedAt:        time.Now().UTC(),
		Title:            "Task 001 updated",
		Description:      "Task 001 description updated",
		Status:           "PAUSED",
		CompleteDeadline: time.Now().UTC().Add(7 * 24 * time.Hour),
		UserID:           actual.UserID,
	}

	if !areTasksEqual(expected, actual) {
		t.Errorf("Response body content error, expected: %s, actual: %s", expected, actual)
	}
}

// Trying to update non-existing task, with taskId: 012d38ae-f234-46f8-bea8-4a3f43b0abb1
func TestUpdateTaskFail(t *testing.T) {
	p := os.Getenv("PORT")

	idStr := "012d38ae-f234-46f8-bea8-4a3f43b0abb1"
	completeDeadline := time.Now().UTC().Add(7 * 24 * time.Hour).Format("2006-01-02T15:04:05.999999Z")
	body := domain.UpdateTaskParams{
		Title:            "Task 001 updated twice",
		Description:      "Task 001 description updated twice",
		Status:           "ACTIVE",
		CompleteDeadline: completeDeadline,
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		t.Fatalf("Could not marshal body: %v", err)
	}

	req, err := http.NewRequest("PUT",
		fmt.Sprintf("http://127.0.0.1:%s/todo/tasks/%s", p, idStr),
		bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatalf("Creating PUT request error: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		t.Fatalf("Executing PUT request error: %v", err)
	}

	defer func(body io.ReadCloser) {
		err := body.Close()
		if err != nil {
			t.Errorf("Closing response body error: %v", err)
		}
	}(res.Body)

	if http.StatusInternalServerError != res.StatusCode {
		t.Errorf("Expected status: %d, actual: %d", http.StatusInternalServerError, res.StatusCode)
	}
}

func TestDeleteTask(t *testing.T) {
	p := os.Getenv("PORT")

	taskId := "012d38ae-f234-46f8-bea8-4a3f43b0abb0"
	req, err := http.NewRequest("DELETE",
		fmt.Sprintf("http://127.0.0.1:%s/todo/tasks/%s", p, taskId),
		nil)
	if err != nil {
		t.Fatalf("Creating DELETE request error: %v", err)
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		t.Fatalf("Executing DELETE request error: %v", err)
	}

	if http.StatusOK != res.StatusCode {
		t.Errorf("Expected status: %d, actual: %d", http.StatusOK, res.StatusCode)
	}
}

func areTaskSlicesEqual(expected []domain.Task, actual []domain.Task) bool {
	if len(expected) != len(actual) {
		return false
	}

	for i, _ := range expected {
		if !areTasksEqual(expected[i], actual[i]) {
			return false
		}
	}

	return true
}

func areTasksEqual(expected domain.Task, actual domain.Task) bool {
	expected.CreatedAt = expected.CreatedAt.Add(5 * time.Minute)
	expected.UpdatedAt = expected.UpdatedAt.Add(5 * time.Minute)
	expected.CompleteDeadline = expected.CompleteDeadline.Add(5 * time.Minute)

	return actual.ID == expected.ID &&
		expected.CreatedAt.After(actual.CreatedAt) &&
		expected.UpdatedAt.After(actual.UpdatedAt) &&
		expected.Title == actual.Title &&
		expected.Description == actual.Description &&
		expected.Status == actual.Status &&
		expected.CompleteDeadline.After(actual.CompleteDeadline) &&
		expected.UserID == actual.UserID
}
