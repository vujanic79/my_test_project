package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/vujanic79/golang-react-todo-app/pkg/domain"
	"io"
	"net/http"
	"os"
	"slices"
	"testing"
)

func TestGetTaskStatuses(t *testing.T) {
	p := os.Getenv("PORT")

	req, err := http.NewRequest("GET",
		fmt.Sprintf("http://127.0.0.1:%s/todo/task-status", p), nil)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		t.Fatalf("Executing GET request error: %v", err)
	}

	defer func(body io.ReadCloser) {
		err := body.Close()
		if err != nil {
			t.Errorf("Closing response body error: %v", err)
		}
	}(res.Body)

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status: %d, actual: %d", http.StatusOK, res.StatusCode)
	}

	decoder := json.NewDecoder(res.Body)
	var actual []domain.TaskStatus
	err = decoder.Decode(&actual)
	if err != nil {
		t.Fatalf("Decoding response body error: %v", err)
	}

	var expected = []domain.TaskStatus{
		{Status: "ACTIVE"},
		{Status: "PAUSED"},
		{Status: "COMPLETED"},
		{Status: "DELETED"}}

	if !slices.Equal(expected, actual) {
		t.Errorf("Response body content error, actual: %s, expected: %s", actual, expected)
	}
}

func TestGetTaskStatusByStatusPass(t *testing.T) {
	p := os.Getenv("PORT")

	req, err := http.NewRequest("GET",
		fmt.Sprintf("http://127.0.0.1:%s/todo/task-status/%s", p, "ACTIVE"), nil)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		t.Fatalf("Executing GET request error: %v", err)
	}

	defer func(body io.ReadCloser) {
		err := body.Close()
		if err != nil {
			t.Errorf("Closing response body error: %v", err)
		}
	}(res.Body)

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status: %d, actual: %d", http.StatusOK, res.StatusCode)
	}

	decoder := json.NewDecoder(res.Body)
	var actual domain.TaskStatus
	err = decoder.Decode(&actual)
	if err != nil {
		t.Fatalf("Decoding response body error: %v", err)
	}

	var expected = domain.TaskStatus{Status: "ACTIVE"}

	if actual != expected {
		t.Errorf("Response body content error, actual: %s, expected: %s", actual, expected)
	}
}

// Trying to get non-existing task status: "ACTIVIE"
func TestGetTaskStatusByStatusFail(t *testing.T) {
	p := os.Getenv("PORT")

	req, err := http.NewRequest("GET",
		fmt.Sprintf("http://127.0.0.1:%s/todo/task-status/%s", p, "ACTIVIE"), nil)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		t.Fatalf("Executing GET request error: %v", err)
	}

	defer func(body io.ReadCloser) {
		err := body.Close()
		if err != nil {
			t.Errorf("Closing response body error: %v", err)
		}
	}(res.Body)

	if res.StatusCode != http.StatusInternalServerError {
		t.Errorf("Expected status: %d, actual: %d", http.StatusInternalServerError, res.StatusCode)
	}
}

func TestCreateTaskStatusPass(t *testing.T) {
	p := os.Getenv("PORT")

	body := domain.CreateTaskStatusParams{
		Status: "ACTIVIE",
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		t.Fatalf("Could not marshal body: %v", err)
	}

	req, err := http.NewRequest("POST",
		fmt.Sprintf("http://127.0.0.1:%s/todo/task-status", p),
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

	if res.StatusCode != http.StatusCreated {
		t.Errorf("Expected status: %d, actual: %d", http.StatusCreated, res.StatusCode)
	}

	decoder := json.NewDecoder(res.Body)
	var actual domain.TaskStatus
	err = decoder.Decode(&actual)
	if err != nil {
		t.Fatalf("Decoding response body error: %v", err)
	}

	var expected = domain.TaskStatus{Status: "ACTIVIE"}

	if actual != expected {
		t.Errorf("Response body content error, actual: %s, expected: %s", actual, expected)
	}
}

// Trying to create already existing status: "ACTIVE"
func TestCreateTaskStatusFail(t *testing.T) {
	p := os.Getenv("PORT")

	body := domain.CreateTaskStatusParams{
		Status: "ACTIVE",
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		t.Fatalf("Could not marshal body: %v", err)
	}

	req, err := http.NewRequest("POST",
		fmt.Sprintf("http://127.0.0.1:%s/todo/task-status", p),
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

	if res.StatusCode != http.StatusInternalServerError {
		t.Errorf("Expected status: %d, actual: %d", http.StatusInternalServerError, res.StatusCode)
	}
}
