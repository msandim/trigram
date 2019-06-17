package server

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/msandim/trigram/store"
)

func init() {
	// Initiate the server:
	server := NewServer(&TestStore{}, 8081)
	go server.Run()
}

func TestLearn(t *testing.T) {
	tests := []struct {
		contentType       string
		text              string
		responseErrorCode int
		responseBody      string
	}{
		{"application/json", "Hello, how are you?", http.StatusMethodNotAllowed, "/learn only accepts POST requests with 'text/plain' as 'Content-type'\n"},
		{"text/plain", "Hello, how are you?", http.StatusOK, ""},
		{"text/plain", "hey", http.StatusBadRequest, "error while learning text: text to learn needs to have more than 3 words\n"},
	}

	for _, test := range tests {
		resp, err := http.Post("http://localhost:8081/learn", test.contentType, bytes.NewBuffer([]byte(test.text)))

		if err != nil {
			t.Fatalf("Unexpected error while doing request: %s", err)
		}

		if resp.StatusCode != test.responseErrorCode {
			t.Fatalf("Unexpected error code: %d", resp.StatusCode)
		}

		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			t.Fatalf("Unexpected error while reading response body: %s", err)
		}

		if string(body) != test.responseBody {
			t.Fatalf("Unexpected body response: \"%s\"", string(body))
		}
	}

	// Additional test for when this method is called as GET:
	resp, err := http.Get("http://localhost:8081/learn")

	if err != nil {
		t.Fatalf("Unexpected error while doing request: %s", err)
	}

	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Fatalf("Unexpected error code received: %d", resp.StatusCode)
	}
}

func TestGenerate(t *testing.T) {

	// Successful test:
	resp, err := http.Get("http://localhost:8081/generate")

	if err != nil {
		t.Fatalf("Unexpected error while doing request: %s", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Fatalf("Unexpected error while reading body: %s", err)
	}

	if string(body) != "Hello and goodbye" {
		t.Fatalf("Unexpected body: %s", string(body))
	}

	// Additional test for when this method is called as POST:
	resp, err = http.Post("http://localhost:8081/generate", "text/plain", bytes.NewBuffer([]byte("")))

	if err != nil {
		t.Fatalf("Unexpected error while doing request: %s", err)
	}

	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Fatalf("Unexpected error code received: %d", resp.StatusCode)
	}
}

type TestStore struct{}

func (store *TestStore) AddTrigram(trigram store.Trigram) {
	fmt.Println(trigram)
}

func (store *TestStore) MakeText() string {
	return "Hello and goodbye"
}
