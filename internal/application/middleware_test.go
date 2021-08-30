package application

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
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
		expectedBody   *FailureResponse
	}{
		{
			name:           "test_limiter",
			method:         http.MethodGet,
			route:          "/api/v1",
			expectedStatus: http.StatusTooManyRequests,
			expectedBody:   NewFailureResponse(http.StatusTooManyRequests, "You have performed too many requests! Please try again in a minute!"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			for i := 0; i < limit; i++ {
				r := httptest.NewRequest(test.method, test.route, nil)
				w := httptest.NewRecorder()
				handler.ServeHTTP(w, r)

				if i == limit {
					assert.NotNil(t, w.Body)
					assert.Equal(t, test.expectedStatus, w.Code)
					assert.JSONEq(t, structToJSON(test.expectedBody), w.Body.String())
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

		assert.Equal(t, http.StatusPermanentRedirect, w.Code)
		assert.Equal(t, "https://example.com/api/v1", w.Result().Header.Get("Location"))
	})
}
