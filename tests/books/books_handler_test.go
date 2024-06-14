package books_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"studies-1/internals/books"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateBookHandler(t *testing.T) {
	// Create a mock request body
	requestBody := []byte(`{"id": "1", "title": "Test Book", "author": "Test Author", "price": 10.99}`)

	// Create a mock HTTP request with the request body
	req, err := http.NewRequest("POST", "/api/v1/books", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	// Create a mock HTTP response recorder
	rr := httptest.NewRecorder()

	// Call the CreateBookHandler function with the mock request and response
	books.CreateBookHandler(rr, req)

	// Check the status code of the response
	assert.Equal(t, http.StatusCreated, rr.Code, "status code should be 201")

	// Trim any trailing whitespace characters from the actual response body
	actualResponseBody := strings.TrimSpace(rr.Body.String())

	// Check the response body
	expectedResponseBody := `{"id":"1","title":"Test Book","author":"Test Author","price":10.99}`
	assert.Equal(t, expectedResponseBody, actualResponseBody, "response body should match")
}

func TestCreateBookHandler_InvalidJSON(t *testing.T) {
	// Create a mock request with invalid JSON
	req, err := http.NewRequest("POST", "/api/v1/books", bytes.NewBuffer([]byte(`invalid-json`)))
	if err != nil {
		t.Fatal(err)
	}

	// Create a mock HTTP response recorder
	rr := httptest.NewRecorder()

	// Call the CreateBookHandler function with the mock request and response
	books.CreateBookHandler(rr, req)

	// Check the status code of the response
	assert.Equal(t, http.StatusBadRequest, rr.Code, "status code should be 400")
}
