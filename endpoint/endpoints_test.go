// endpoints_test.go
package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHealthCheckHandler(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "/health", nil)
	/*if err != nil {
		t.Fatal(err)
	}*/
	/*if ret := assert.Nil(t, err); ret == false {
		return
	}*/
	require.Nil(t, err)

	handler := http.HandlerFunc(HealthCheckHandler)

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()

	fmt.Println("Before ServeHTTP")
	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)
	fmt.Println("After ServeHTTP")

	// Check the status code is what we expect.
	/*if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}*/
	assert.Equalf(t, int(http.StatusOK), rr.Code, "handler returned wrong status code: got %v want %v", rr.Code, http.StatusOK)

	// Check the response body is what we expect.
	expected := `{"alive": true}`
	/*if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}*/
	assert.Equalf(t, expected, rr.Body.String(), "handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
}
