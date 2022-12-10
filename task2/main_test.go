package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

// The TestUpperCaseHandler function tests the upperCaseHandler function.
// It creates a new request and response recorder, calls the handler, and then checks the response.
// If the response is not what we expect, the test fails.
func TestUpperCaseHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/upper?word=hello", nil)
	w := httptest.NewRecorder()
	upperCaseHandler(w, req)

	res := w.Result()

	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}

	if string(data) != "HELLO" {
		t.Errorf("expected HELLO, got %v", string(data))
	}
}
