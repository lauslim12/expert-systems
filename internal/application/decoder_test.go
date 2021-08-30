package application

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJSONDecoder(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	defer ts.Close()

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
		{
			name:           "test_success",
			method:         http.MethodPost,
			input:          `{"name":"Mai Sakurajima","address":"Fujisawa"}`,
			expectedStatus: http.StatusOK,
			expectedBody:   nil,
			withHeader:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequest(tt.method, "/api/v1", strings.NewReader(tt.input))
			w := httptest.NewRecorder()
			if tt.withHeader {
				r.Header.Set("Content-Type", "application/json")
			}

			failureResponse := decodeJSONBody(w, r, &Person{})
			assert.JSONEq(t, structToJSON(tt.expectedBody), structToJSON(failureResponse))
		})
	}
}
