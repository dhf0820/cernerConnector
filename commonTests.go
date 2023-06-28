package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	fmt.Printf("executeRequest:11  -- starting req url: %s\n", req.URL)
	rr := httptest.NewRecorder()
	NewRouter().ServeHTTP(rr, req)
	fmt.Printf("executeRequest:14  -- returning\n")
	return rr
}
func checkResponseCode(t *testing.T, expected, actual int) {
	fmt.Printf("checkResponseCode:18  -- starting  actual= %d  expected= %d\n", actual, expected)
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
		return
	}
	fmt.Printf("checkResponseCode:22  -- All Ok\n")
}
