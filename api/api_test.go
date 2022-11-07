package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type exampleResponse struct {
	Message  string `json:"message"`
	Response string `json:"response"`
}

func TestRedirectUrl(t *testing.T) {
	router := SetupRouter()
	ctx := exampleResponse{Message: "pong", Response: "ping"}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/:url", nil)
	router.ServeHTTP(w, req)

	wantBody, _ := json.Marshal(ctx)

	assertCode(t, http.StatusOK, w.Code)
	assertBody(t, string(wantBody), w.Body.String())
}

func assertCode(t testing.TB, want, got int) {
	t.Helper()

	if want != got {
		t.Fatalf("got %d code, but want %d code", got, want)
	}
}

func assertBody(t testing.TB, want, got string) {
	t.Helper()

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("got %+v body, but want %+v body", got, want)
	}
}
