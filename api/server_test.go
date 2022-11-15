package api

// type exampleResponse struct {
// 	Message  string `json:"message"`
// 	Response string `json:"response"`
// }

// type mockResponse gin.H

// func TestAllRoute(t *testing.T) {
// 	router := SetupRouter()

// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("GET", "/api/all", nil)
// 	router.ServeHTTP(w, req)

// 	wantBody := gin.H{
// 		"ggl": "https://google.com",
// 		"go":  "https://go.dev",
// 		"tw":  "https://twitch.tv",
// 	}
// 	body := gin.H{}
// 	_ = json.Unmarshal(w.Body.Bytes(), &body)

// 	assertStatusCode(t, http.StatusOK, w.Code)
// 	assertResponseBody(t, wantBody, body)
// }

// func assertStatusCode(t testing.TB, want, got int) {
// 	t.Helper()

// 	if want != got {
// 		t.Fatalf("got %d code, but want %d code", got, want)
// 	}
// }

// func assertResponseBody(t testing.TB, want, got gin.H) {
// 	t.Helper()

// 	if !reflect.DeepEqual(want, got) {
// 		t.Fatalf("got %+v body, but want %+v body", got, want)
// 	}
// }
