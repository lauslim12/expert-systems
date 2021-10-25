package application

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLimiterHandler(t *testing.T) {
	limit := 250
	handler := Configure(pathToWebDirectory, applicationModeDevelopment)
	testServer := httptest.NewServer(handler)
	defer testServer.Close()

	// Define test-cases.
	tests := []struct {
		name           string
		method         string
		route          string
		expectedStatus int
	}{
		{
			name:           "test_limiter",
			method:         http.MethodGet,
			route:          "/api/v1",
			expectedStatus: http.StatusTooManyRequests,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			for i := 0; i < limit; i++ {
				r := httptest.NewRequest(test.method, test.route, nil)
				w := httptest.NewRecorder()
				handler.ServeHTTP(w, r)

				if i == limit {
					if test.expectedStatus != w.Code {
						t.Errorf("Expected and actual status code values are different! Expected: %v. Got: %v", test.expectedStatus, w.Code)
					}
				}
			}
		})
	}
}

func TestHTTPSRedirectOnProduction(t *testing.T) {
	handler := Configure(pathToWebDirectory, applicationModeProduction)
	testServer := httptest.NewServer(handler)
	defer testServer.Close()

	t.Run("test_https_redirect_on_production", func(t *testing.T) {
		r := httptest.NewRequest("GET", "/api/v1", nil)
		w := httptest.NewRecorder()
		r.Header.Add("X-Forwarded-Proto", "http")
		handler.ServeHTTP(w, r)

		if http.StatusPermanentRedirect != w.Code {
			t.Errorf("Expected and actual status code values are different! Expected: %v. Got: %v", http.StatusMovedPermanently, w.Code)
		}

		if w.Result().Header.Get("Location") != "https://example.com/api/v1" {
			t.Errorf("Expected and actual location values are different! Expected: %v. Got: %v", "https://example.com/api/v1", w.Result().Header.Get("Location"))
		}
	})
}
