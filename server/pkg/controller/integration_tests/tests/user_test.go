package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/vujanic79/golang-react-todo-app/pkg/domain"
	"io"
	"net/http"
	"os"
	"testing"
	"time"
)

func TestCreateUserPass(t *testing.T) {
	p := os.Getenv("PORT")

	body := domain.CreateUserParams{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "johndoe@gmail.com",
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		t.Fatalf("Could not marshal body: %v", err)
	}

	req, err := http.NewRequest("POST",
		fmt.Sprintf("http://127.0.0.1:%s/todo/users", p),
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
	var actual domain.User
	err = decoder.Decode(&actual)
	if err != nil {
		t.Fatalf("Decoding response body error: %v", err)
	}

	var expected = domain.User{
		ID:        actual.ID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		FirstName: "John",
		LastName:  "Doe",
		Email:     "johndoe@gmail.com",
	}

	if !areUsersEqual(expected, actual) {
		t.Errorf("Response body content error, expected: %s, actual: %s", expected, actual)
	}
}

// Trying to create user with existing mail: johndoe@gmail.com
func TestCreateUserFail(t *testing.T) {
	p := os.Getenv("PORT")

	body := domain.CreateUserParams{
		FirstName: "Milan",
		LastName:  "Vujanic",
		Email:     "milanmvujanic@gmail.com",
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		t.Fatalf("Could not marshal body: %v", err)
	}

	req, err := http.NewRequest("POST",
		fmt.Sprintf("http://127.0.0.1:%s/todo/users", p),
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

func areUsersEqual(expected domain.User, actual domain.User) bool {
	expected.CreatedAt = expected.CreatedAt.Add(5 * time.Minute)
	expected.UpdatedAt = expected.UpdatedAt.Add(5 * time.Minute)

	return expected.ID == actual.ID &&
		expected.CreatedAt.After(actual.CreatedAt) &&
		expected.UpdatedAt.After(actual.UpdatedAt) &&
		expected.FirstName == actual.FirstName &&
		expected.LastName == actual.LastName &&
		expected.Email == actual.Email
}
