package application

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Variable 'pathtoWebDirectory' is intentionally kept to './web' and not './web/build' for the sake of testing.
const pathToWebDirectory = "./web"

func structToJSON(object interface{}) string {
	out, err := json.Marshal(object)
	if err != nil {
		log.Fatal(err.Error())
	}

	return string(out)
}

func TestGeneralHandler(t *testing.T) {
	handler := Configure(pathToWebDirectory, applicationModeDevelopment)
	testServer := httptest.NewServer(handler)
	defer testServer.Close()

	failureTests := []struct {
		name           string
		method         string
		input          string
		route          string
		expectedStatus int
	}{
		{
			name:           "test_method_not_allowed",
			method:         http.MethodPut,
			input:          "",
			route:          "/api/v1",
			expectedStatus: http.StatusMethodNotAllowed,
		},
		{
			name:           "test_route_not_exist",
			method:         http.MethodGet,
			input:          "",
			route:          "/api/v1/404",
			expectedStatus: http.StatusNotFound,
		},
		{
			name:           "test_fail_endpoint",
			method:         http.MethodPost,
			input:          "",
			route:          "/api/v1",
			expectedStatus: http.StatusUnsupportedMediaType,
		},
	}

	successTests := []struct {
		name           string
		method         string
		input          string
		route          string
		expectedStatus int
	}{
		{
			name:           "test_health",
			method:         http.MethodGet,
			input:          "",
			route:          "/api/v1",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range failureTests {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequest(tt.method, tt.route, nil)
			w := httptest.NewRecorder()
			handler.ServeHTTP(w, r)

			assert.Equal(t, tt.expectedStatus, w.Code)
		})
	}

	for _, tt := range successTests {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequest(tt.method, tt.route, strings.NewReader(tt.input))
			w := httptest.NewRecorder()
			r.Header.Set("Content-Type", "application/json")
			handler.ServeHTTP(w, r)

			assert.Equal(t, tt.expectedStatus, w.Code)
		})
	}
}

// Create a custom recorder so we can read from static files.
// Reference: https://github.com/go-chi/chi/issues/583.
type testRecorder struct {
	*httptest.ResponseRecorder
}

func (rec *testRecorder) ReadFrom(r io.Reader) (n int64, err error) {
	return io.Copy(rec.ResponseRecorder, r)
}

func newRecorder() *testRecorder {
	return &testRecorder{ResponseRecorder: httptest.NewRecorder()}
}

func TestRenderWeb(t *testing.T) {
	handler := Configure(pathToWebDirectory, applicationModeDevelopment)
	testServer := httptest.NewServer(handler)
	defer testServer.Close()

	// Reverse current working directory to the root folder.
	// This is done so the test can reach the 'pathToWebDirectory' location.
	err := os.Chdir(filepath.Join("..", ".."))
	if err != nil {
		log.Fatal(err)
	}

	t.Run("test_render_web", func(t *testing.T) {
		r := httptest.NewRequest(http.MethodGet, "/", nil)
		w := newRecorder()
		handler.ServeHTTP(w, r)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("test_render_web_404", func(t *testing.T) {
		r := httptest.NewRequest(http.MethodGet, "/404", nil)
		w := newRecorder()
		handler.ServeHTTP(w, r)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})
}
