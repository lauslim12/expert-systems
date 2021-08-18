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

// This variable is intentionally kept to './web' and not './web/build' for the sake of testing.
const pathToWebDirectory = "./web"

func TestSuccessGeneralHandler(t *testing.T) {
	handler := Configure(pathToWebDirectory)
	testServer := httptest.NewServer(handler)
	defer testServer.Close()

	// Define test-cases.
	tests := []struct {
		name           string
		method         string
		route          string
		expectedStatus int
		expectedBody   *SuccessResponse
	}{
		{
			name:           "test_health",
			method:         http.MethodGet,
			route:          "/api/v1",
			expectedStatus: http.StatusOK,
			expectedBody:   NewSuccessResponse(http.StatusOK, "Welcome to 'net/http' API!", nil),
		},
	}

	// Perform requests.
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := httptest.NewRequest(test.method, test.route, nil)
			w := httptest.NewRecorder()
			handler.ServeHTTP(w, r)

			out, err := json.Marshal(test.expectedBody)
			if err != nil {
				log.Fatal(err.Error())
			}

			assert.NotNil(t, w.Body)
			assert.Equal(t, test.expectedStatus, w.Code)
			assert.JSONEq(t, string(out), w.Body.String())
		})
	}
}

func TestFailureGeneralHandler(t *testing.T) {
	handler := Configure(pathToWebDirectory)
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
			name:           "test_method_not_allowed",
			method:         http.MethodPut,
			route:          "/api/v1",
			expectedStatus: http.StatusMethodNotAllowed,
			expectedBody:   NewFailureResponse(http.StatusMethodNotAllowed, "Method 'PUT' is not allowed in this route!"),
		},
		{
			name:           "test_route_not_exist",
			method:         http.MethodGet,
			route:          "/api/v1/404",
			expectedStatus: http.StatusNotFound,
			expectedBody:   NewFailureResponse(http.StatusNotFound, "Route '/api/v1/404' does not exist in this server!"),
		},
	}

	// Perform requests.
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := httptest.NewRequest(test.method, test.route, nil)
			w := httptest.NewRecorder()
			handler.ServeHTTP(w, r)

			out, err := json.Marshal(test.expectedBody)
			if err != nil {
				log.Fatal(err.Error())
			}

			assert.NotNil(t, w.Body)
			assert.Equal(t, test.expectedStatus, w.Code)
			assert.JSONEq(t, string(out), w.Body.String())
		})
	}
}

func TestLimiterHandler(t *testing.T) {
	limit := 250
	handler := Configure(pathToWebDirectory)
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
					out, err := json.Marshal(test.expectedBody)
					if err != nil {
						log.Fatal(err.Error())
					}

					assert.Equal(t, test.expectedStatus, w.Code)
					assert.NotNil(t, w.Body)
					assert.JSONEq(t, string(out), w.Body.String())
				}
			}
		})
	}
}

func TestSuccessFunctionalHandler(t *testing.T) {
	handler := Configure(pathToWebDirectory)
	testServer := httptest.NewServer(handler)
	defer testServer.Close()

	// Create sample datasets.
	kaede := &Person{
		Name:    "Kaede Kimura",
		Address: "Tokyo",
	}

	// Define test cases.
	tests := []struct {
		name           string
		method         string
		input          string
		expectedStatus int
		expectedBody   *SuccessResponse
	}{
		{
			name:           "test_success",
			method:         http.MethodPost,
			input:          `{"name":"Kaede Kimura","address":"Tokyo"}`,
			expectedStatus: http.StatusOK,
			expectedBody:   NewSuccessResponse(http.StatusOK, "Successfully processed data in the Expert Systems!", kaede),
		},
	}

	// Perform requests.
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := httptest.NewRequest(test.method, "/api/v1", strings.NewReader(test.input))
			r.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			handler.ServeHTTP(w, r)

			out, err := json.Marshal(test.expectedBody)
			if err != nil {
				log.Fatal(err.Error())
			}

			assert.NotNil(t, w.Body)
			assert.Equal(t, test.expectedStatus, w.Code)
			assert.JSONEq(t, string(out), w.Body.String())
		})
	}
}

func TestFailureFunctionalHandler(t *testing.T) {
	handler := Configure(pathToWebDirectory)
	testServer := httptest.NewServer(handler)
	defer testServer.Close()

	// Define test-cases.
	tests := []struct {
		name           string
		method         string
		input          string
		expectedStatus int
		expectedBody   *FailureResponse
		withHeader     bool
	}{
		{
			name:           "test_failure_no_header",
			method:         http.MethodPost,
			input:          `{"name":"Kaede Kimura"}`,
			expectedStatus: http.StatusUnsupportedMediaType,
			expectedBody:   NewFailureResponse(http.StatusUnsupportedMediaType, "The 'Content-Type' header is not 'application/json'!"),
			withHeader:     false,
		},
		{
			name:           "test_failure_empty",
			method:         http.MethodPost,
			input:          "",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   NewFailureResponse(http.StatusBadRequest, "Request body must not be empty!"),
			withHeader:     true,
		},
		{
			name:           "test_failure_not_validated",
			method:         http.MethodPost,
			input:          `{"name":"Kaede Kimura"}`,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   NewFailureResponse(http.StatusBadRequest, "Key: 'Person.Address' Error:Field validation for 'Address' failed on the 'required' tag"),
			withHeader:     true,
		},
		{
			name:           "test_failure_bad_format_json_position",
			method:         http.MethodPost,
			input:          `{"name":"Kaede Kimura","address":"Kyoto",badfomathere}`,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   NewFailureResponse(http.StatusBadRequest, "Request body contains a badly formatted JSON at position 42!"),
			withHeader:     true,
		},
		{
			name:           "test_failure_invalid_value",
			method:         http.MethodPost,
			input:          `{"name":"Kaede Kimura","address":12345}`,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   NewFailureResponse(http.StatusBadRequest, "Request body contains an invalid value for the \"address\" field at position 38!"),
			withHeader:     true,
		},
		{
			name:           "test_failure_bad_json",
			method:         http.MethodPost,
			input:          `{"name":"Kaede Kimura","address":12345`,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   NewFailureResponse(http.StatusBadRequest, "Request body contains a badly-formed JSON!"),
			withHeader:     true,
		},
		{
			name:           "test_failure_array_json",
			method:         http.MethodPost,
			input:          `[{"name":"Kaede Kimura","address":"12345"},{"name":"Mai Sakurajima","address":"Fujisawa"}]`,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   NewFailureResponse(http.StatusBadRequest, "Request body contains an invalid value for the \"\" field at position 1!"),
			withHeader:     true,
		},
		{
			name:           "test_failure_unknown_fields",
			method:         http.MethodPost,
			input:          `{"name":"Mai Sakurajima","mockAttribute":"Fujisawa"}`,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   NewFailureResponse(http.StatusBadRequest, "Request body contains unknown field '\"mockAttribute\"'!"),
			withHeader:     true,
		},
		{
			name:           "test_failure_single_json",
			method:         http.MethodPost,
			input:          `{"name":"Mai Sakurajima","address":"Fujisawa"}{"name":"Kamisato Ayaka","address":"Fukuoka"}`,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   NewFailureResponse(http.StatusBadRequest, "Request body must only contain a single JSON object!"),
			withHeader:     true,
		},
		{
			name:           "test_failure_payload_size",
			method:         http.MethodPost,
			input:          `{"name":"Mai Sakurajima","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa"}`,
			expectedStatus: http.StatusRequestEntityTooLarge,
			expectedBody:   NewFailureResponse(http.StatusRequestEntityTooLarge, "Request body must not be larger than 1KB!"),
			withHeader:     true,
		},
	}

	// Perform requests.
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := httptest.NewRequest(test.method, "/api/v1", strings.NewReader(test.input))
			w := httptest.NewRecorder()

			if test.withHeader {
				r.Header.Set("Content-Type", "application/json")
			}

			handler.ServeHTTP(w, r)

			out, err := json.Marshal(test.expectedBody)
			if err != nil {
				log.Fatal(err.Error())
			}

			assert.NotNil(t, w.Body)
			assert.Equal(t, test.expectedStatus, w.Code)
			assert.JSONEq(t, string(out), w.Body.String())
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
	handler := Configure(pathToWebDirectory)
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
