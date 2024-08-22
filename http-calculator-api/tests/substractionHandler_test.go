package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sabir222/http-calculator/internal/handler"
	"testing"
)

type Result struct {
	Result int `json:"Result"`
}

func TestSubstractionHandler(t *testing.T) {

	server := httptest.NewServer(http.HandlerFunc(handler.SubstractionHandler))
	defer server.Close()

	reqBody := map[string][]int{
		"Numbers": {10, 2, 3},
	}

	body, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := http.Post(server.URL, "application/json", bytes.NewBuffer(body))

	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200, but got %d", resp.StatusCode)
	}

	result := Result{}
	expected := 5

	decoder := json.NewDecoder(resp.Body)

	if err := decoder.Decode(&result); err != nil {
		t.Fatal(err)
	}

	if result.Result != expected {
		t.Errorf("Expected %d but we got %d", expected, result.Result)
	}

}
