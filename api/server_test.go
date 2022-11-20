package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Patrick564/url-shortener-backend/api/controllers"
	"github.com/Patrick564/url-shortener-backend/internal/models"
	"github.com/Patrick564/url-shortener-backend/utils"
)

type mockUrlResponse struct {
	Error error        `json:"error"`
	Url   *models.Url  `json:"url,omitempty"`
	Urls  []models.Url `json:"urls,omitempty"`
}

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

func TestAllRoute(t *testing.T) {
	env := &controllers.Env{Urls: &mockUrlModel{}}
	router := SetupRouter(env)

	mockResponse := mockUrlResponse{
		Error: nil,
		Urls: []models.Url{
			{ShortUrl: "ID_1", OriginalUrl: "https://www.example-url-1.com"},
			{ShortUrl: "ID_2", OriginalUrl: "https://www.example-url-2.dev"},
			{ShortUrl: "ID_3", OriginalUrl: "https://www.example-url-3.com"},
		},
	}
	want, _ := json.Marshal(mockResponse)

	tests := []struct {
		name      string
		wantCode  int
		wantError string
		wantBody  []byte
	}{
		{
			name:      "Returns without errors",
			wantCode:  http.StatusOK,
			wantError: "",
			wantBody:  want,
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

	type m struct {
		Error error      `json:"error"`
		Url   models.Url `json:"url"`
	}

	tests := []struct {
		name            string
		reqBody         io.Reader
		wantCode        int
		wantError       error
		rawMockResponse m
	}{
		{
			name:      "Returns without errors",
			reqBody:   bytes.NewBuffer([]byte("{\"url\": \"https://www.example.com\" }")),
			wantCode:  http.StatusOK,
			wantError: nil,
			rawMockResponse: m{
				Error: nil,
				Url:   models.Url{ShortUrl: "ID_1", OriginalUrl: "https://www.example.com"},
			},
		},
		{
			name:      "Returns error with empty body",
			reqBody:   bytes.NewBuffer([]byte("")),
			wantCode:  http.StatusBadRequest,
			wantError: utils.ErrEmptyBody,
			rawMockResponse: m{
				Error: utils.ErrEmptyBody,
				Url:   models.Url{},
			},
		},
		{
			name:      "Returns error with incorrect url",
			reqBody:   bytes.NewBuffer([]byte("{\"url\": \"ejemplo-bad-url:4040\" }")),
			wantCode:  http.StatusBadRequest,
			wantError: utils.ErrInvalidUrl,
			rawMockResponse: m{
				Error: utils.ErrInvalidUrl,
				Url:   models.Url{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/api/add", tt.reqBody)
			req.Header.Set("Content-Type", "application/json")
			router.ServeHTTP(w, req)

			wantBody, _ := json.Marshal(tt.rawMockResponse)

			utils.AssertError(t, tt.wantError, tt.rawMockResponse.Error)
			utils.AssertStatusCode(t, tt.wantCode, w.Code)
			utils.AssertResponseBody(t, wantBody, w.Body.String())
		})
	}
}
