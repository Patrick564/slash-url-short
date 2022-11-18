package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/Patrick564/url-shortener-backend/api/controllers"
	"github.com/Patrick564/url-shortener-backend/internal/models"
)

type mockUrlResponse struct {
	Error string       `json:"error"`
	Urls  []models.Url `json:"urls"`
}

type mockUrlModel struct{}

func (m *mockUrlModel) GetAll() ([]models.Url, error) {
	var urls []models.Url

	urls = append(urls, models.Url{ShortUrl: "ID_1", OriginalUrl: "www.example-url-1.com"})
	urls = append(urls, models.Url{ShortUrl: "ID_2", OriginalUrl: "www.example-url-2.dev"})
	urls = append(urls, models.Url{ShortUrl: "ID_3", OriginalUrl: "www.example-url-3.com"})

	return urls, nil
}

func TestAllRoute(t *testing.T) {
	env := &controllers.Env{Urls: &mockUrlModel{}}

	router := SetupRouter(env)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/all", nil)
	router.ServeHTTP(w, req)

	mockResponse := mockUrlResponse{
		Error: "",
		Urls: []models.Url{
			{ShortUrl: "ID_1", OriginalUrl: "www.example-url-1.com"},
			{ShortUrl: "ID_2", OriginalUrl: "www.example-url-2.dev"},
			{ShortUrl: "ID_3", OriginalUrl: "www.example-url-3.com"},
		},
	}
	want, _ := json.Marshal(mockResponse)

	assertStatusCode(t, http.StatusOK, w.Code)
	assertResponseBody(t, want, w.Body.String())
}

func assertStatusCode(t testing.TB, want, got int) {
	t.Helper()

	if want != got {
		t.Fatalf("got %d code, but want %d code", got, want)
	}
}

func assertResponseBody(t testing.TB, want []byte, got string) {
	t.Helper()

	if !reflect.DeepEqual(string(want), got) {
		t.Fatalf("got %+v body, but want %+v body", got, string(want))
	}
}
