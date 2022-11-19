package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Patrick564/url-shortener-backend/api/controllers"
	"github.com/Patrick564/url-shortener-backend/internal/models"
	"github.com/Patrick564/url-shortener-backend/utils"
)

type mockUrlResponse struct {
	Error string       `json:"error"`
	Url   *models.Url  `json:"url,omitempty"`
	Urls  []models.Url `json:"urls,omitempty"`
}

type mockUrlModel struct{}

func (m *mockUrlModel) GetAll() ([]models.Url, error) {
	var urls []models.Url

	urls = append(urls, models.Url{ShortUrl: "ID_1", OriginalUrl: "www.example-url-1.com"})
	urls = append(urls, models.Url{ShortUrl: "ID_2", OriginalUrl: "www.example-url-2.dev"})
	urls = append(urls, models.Url{ShortUrl: "ID_3", OriginalUrl: "www.example-url-3.com"})

	return urls, nil
}

func (m *mockUrlModel) Add(id string, url string) (*models.Url, error) {
	u := models.Url{ShortUrl: "ID_1", OriginalUrl: "www.example.com"}

	return &u, nil
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

	utils.AssertStatusCode(t, http.StatusOK, w.Code)
	utils.AssertResponseBody(t, want, w.Body.String())
}

func TestAddRoute(t *testing.T) {
	env := &controllers.Env{Urls: &mockUrlModel{}}
	router := SetupRouter(env)
	w := httptest.NewRecorder()

	reqBody := bytes.NewBuffer([]byte("{\"url\": \"www.example.com\" }"))
	req, _ := http.NewRequest("POST", "/api/add", reqBody)
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	mockResponse := mockUrlResponse{
		Error: "",
		Url:   &models.Url{ShortUrl: "ID_1", OriginalUrl: "www.example.com"},
	}
	want, _ := json.Marshal(mockResponse)

	utils.AssertStatusCode(t, http.StatusOK, w.Code)
	utils.AssertResponseBody(t, want, w.Body.String())
}
