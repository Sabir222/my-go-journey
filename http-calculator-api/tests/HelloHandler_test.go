package tests

import (
	"io"
	"net/http"
	"net/http/httptest"
	"sabir222/http-calculator/internal/handler"
	"testing"
)

func TestHelloHandlerRR(t *testing.T) {

	rr := httptest.NewRecorder()

	req, err := http.NewRequest(http.MethodGet, "", nil)
	if err != nil {
		t.Errorf("Error creating request: %v", err)
	}

	handler.HelloWorldHandler(rr, req)
	if rr.Result().StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200, but got %d", rr.Result().StatusCode)
	}

	defer rr.Result().Body.Close()

	expected := "Hello World!"

	body, err := io.ReadAll(rr.Result().Body)
	if err != nil {
		t.Fatal(err)
	}

	if string(body) != expected {
		t.Errorf("Expected %s but we got %s", expected, string(body))
	}
}

func TestHelloHandler(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(handler.HelloWorldHandler))
	resp, err := http.Get(server.URL)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200, but got %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	expected := "Hello World!"

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	if string(body) != expected {
		t.Errorf("Expected %s but we got %s", expected, string(body))
	}
}
