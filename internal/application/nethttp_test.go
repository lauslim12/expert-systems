package application

import (
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// Variable 'pathtoWebDirectory' is intentionally kept to './web' and not './web/build' for the sake of testing.
const pathToWebDirectory = "./web"

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
		{
			name:           "test_input",
			method:         http.MethodPost,
			input:          `{"symptoms":[{"symptomId":"S1","weight":0.2},{"symptomId":"S2","weight":0.2},{"symptomId":"S3","weight":0.2},{"symptomId":"S4","weight":0.4},{"symptomId":"S5","weight":0.2},{"symptomId":"S6","weight":0.4},{"symptomId":"S7","weight":0.8},{"symptomId":"S8","weight":0.2},{"symptomId":"S9","weight":0.2},{"symptomId":"S10","weight":0.4},{"symptomId":"S11","weight":0.2},{"symptomId":"S12","weight":0.2},{"symptomId":"S13","weight":1}]}`,
			route:          "/api/v1",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range failureTests {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequest(tt.method, tt.route, nil)
			w := httptest.NewRecorder()
			handler.ServeHTTP(w, r)

			if tt.expectedStatus != w.Code {
				t.Errorf("Expected and actual status code values are different! Expected: %v. Got: %v", tt.expectedStatus, w.Code)
			}
		})
	}

	for _, tt := range successTests {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequest(tt.method, tt.route, strings.NewReader(tt.input))
			w := httptest.NewRecorder()
			r.Header.Set("Content-Type", "application/json")
			handler.ServeHTTP(w, r)

			if tt.expectedStatus != w.Code {
				t.Errorf("Expected and actual status code values are different! Expected: %v. Got: %v", tt.expectedStatus, w.Code)
			}
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

		if http.StatusOK != w.Code {
			t.Errorf("Expected and actual status code values are different! Expected: %v. Got: %v", http.StatusOK, w.Code)
		}
	})

	t.Run("test_render_web_404", func(t *testing.T) {
		r := httptest.NewRequest(http.MethodGet, "/404", nil)
		w := newRecorder()
		handler.ServeHTTP(w, r)

		if http.StatusNotFound != w.Code {
			t.Errorf("Expected and actual status code values are different! Expected: %v. Got: %v", http.StatusNotFound, w.Code)
		}
	})
}
