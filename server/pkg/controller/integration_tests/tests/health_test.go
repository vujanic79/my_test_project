package tests

import (
	"fmt"
	"net/http"
	"os"
	"testing"
)

func TestHealthRoute(t *testing.T) {
	p := os.Getenv("PORT")
	req, err := http.NewRequest("GET", fmt.Sprintf("http://127.0.0.1:%s/todo/healthz", p), nil)
	if err != nil {
		t.Fatalf("Creating GET request error: %v", err)
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		t.Fatalf("Executing GET request error: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status: %d, actual: %d", http.StatusOK, res.StatusCode)
	}
}
