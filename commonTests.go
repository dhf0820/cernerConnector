package main

import (
	"fmt"
	log "github.com/dhf0820/vslog"
	"net/http"
	"net/http/httptest"
	"testing"
)

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	log.Debug3("executeRequest  -- starting req url: " + fmt.Sprint(req.URL))
	rr := httptest.NewRecorder()
	NewRouter().ServeHTTP(rr, req)
	log.Debug3("executeRequest  -- returning")
	return rr
}
func checkResponseCode(t *testing.T, expected, actual int) {
	log.ErrMsg(fmt.Sprintf("checkResponseCode:  -- starting  actual= %d  expected= %d\n", actual, expected))
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
		return
	}
	log.Info("checkResponseCode: All Ok")
}
