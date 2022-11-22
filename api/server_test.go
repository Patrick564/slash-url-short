package api

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Patrick564/url-shortener-backend/api/controllers"
	"github.com/Patrick564/url-shortener-backend/internal/models"
	"github.com/Patrick564/url-shortener-backend/utils"
)

type mockUrlModel struct{}

func (m *mockUrlModel) All() ([]models.Url, error) {
	var urls []models.Url

	urls = append(urls, models.Url{ShortUrl: "ID_1", OriginalUrl: "https://www.example-url-1.com"})
	urls = append(urls, models.Url{ShortUrl: "ID_2", OriginalUrl: "https://www.example-url-2.dev"})
	urls = append(urls, models.Url{ShortUrl: "ID_3", OriginalUrl: "https://www.example-url-3.com"})

	return urls, nil
}

func (m *mockUrlModel) Add(url string) (models.Url, error) {
	if url != "https://www.example.com" {
		return models.Url{}, utils.ErrEmptyBody
	}

	return models.Url{ShortUrl: "ID_1", OriginalUrl: "https://www.example.com"}, nil
}

func (m *mockUrlModel) GoTo(id string) (string, error) {
	if id == "incorrect-id" {
		return "", utils.ErrInvalidID
	}

	return "https://www.google.com", nil
}

func TestAllRoute(t *testing.T) {
	env := &controllers.Env{Urls: &mockUrlModel{}}
	router := SetupRouter(env)

	tests := []struct {
		name     string
		wantCode int
		wantBody string
	}{
		{
			name:     "Returns without errors",
			wantCode: http.StatusOK,
			wantBody: `{"urls":[{"short_url":"ID_1","original_url":"https://www.example-url-1.com"},{"short_url":"ID_2","original_url":"https://www.example-url-2.dev"},{"short_url":"ID_3","original_url":"https://www.example-url-3.com"}]}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/api/all", nil)
			router.ServeHTTP(w, req)

			utils.AssertStatusCode(t, tt.wantCode, w.Code)
			utils.AssertResponseBody(t, tt.wantBody, w.Body.String())
		})
	}
}

func TestAddRoute(t *testing.T) {
	env := &controllers.Env{Urls: &mockUrlModel{}}
	router := SetupRouter(env)

	tests := []struct {
		name     string
		body     io.Reader
		wantCode int
		wantBody string
	}{
		{
			name:     "Returns without errors",
			body:     bytes.NewBuffer([]byte("{\"url\": \"https://www.example.com\" }")),
			wantCode: http.StatusOK,
			wantBody: `{"url":{"short_url":"ID_1","original_url":"https://www.example.com"}}`,
		},
		{
			name:     "Returns error with empty body",
			body:     bytes.NewBuffer([]byte("")),
			wantCode: http.StatusBadRequest,
			wantBody: `{"error":"empty body"}`,
		},
		{
			name:     "Returns error with incorrect url",
			body:     bytes.NewBuffer([]byte("{\"url\": \"ejemplo-bad-url:4040\" }")),
			wantCode: http.StatusBadRequest,
			wantBody: `{"error":"empty body"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/api/add", tt.body)
			req.Header.Set("Content-Type", "application/json")
			router.ServeHTTP(w, req)

			utils.AssertStatusCode(t, tt.wantCode, w.Code)
			utils.AssertResponseBody(t, tt.wantBody, w.Body.String())
		})
	}
}

func TestRedirectIDRoute(t *testing.T) {
	env := &controllers.Env{Urls: &mockUrlModel{}}
	router := SetupRouter(env)

	tests := []struct {
		name     string
		id       string
		wantCode int
		wantBody string
	}{
		{
			name:     "Redirects without errors",
			id:       "/api/correct-id",
			wantCode: http.StatusMovedPermanently,
			wantBody: "<a href=\"https://www.google.com\">Moved Permanently</a>.\n\n",
		},
		{
			name:     "Redirects with incorrect id error",
			id:       "/api/incorrect-id",
			wantCode: http.StatusBadRequest,
			wantBody: "{\"error\":\"incorrect short url\",\"status\":400}",
		},
		{
			name:     "Redirects with empty id error",
			id:       "/api/",
			wantCode: http.StatusNotFound,
			wantBody: "404 page not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", tt.id, nil)
			router.ServeHTTP(w, req)

			utils.AssertStatusCode(t, tt.wantCode, w.Code)
			utils.AssertResponseBody(t, tt.wantBody, w.Body.String())
		})
	}
}
