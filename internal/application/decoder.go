package application

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/go-playground/validator"
)

// Utility function to decode a JSON request body.
func decodeJSONBody(w http.ResponseWriter, r *http.Request, dst interface{}) *FailureResponse {
	// Check if Header is 'Content-Type: application/json'.
	if r.Header.Get("Content-Type") != "application/json" {
		return NewFailureResponse(http.StatusUnsupportedMediaType, "The 'Content-Type' header is not 'application/json'!")
	}

	// Parse body, and set max bytes reader (1KB).
	// No defaults because I believe there are no more possible errors. Please tell me if you think otherwise!
	r.Body = http.MaxBytesReader(w, r.Body, 1024)
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(dst); err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError

		switch {
		// Handle syntax errors.
		case errors.As(err, &syntaxError):
			errorMessage := fmt.Sprintf("Request body contains a badly formatted JSON at position %d!", syntaxError.Offset)
			return NewFailureResponse(http.StatusBadRequest, errorMessage)

		// Handle unexpected EOFs.
		case errors.Is(err, io.ErrUnexpectedEOF):
			errorMessage := "Request body contains a badly-formed JSON!"
			return NewFailureResponse(http.StatusBadRequest, errorMessage)

		// Handle wrong data-type in request body.
		case errors.As(err, &unmarshalTypeError):
			errorMessage := fmt.Sprintf("Request body contains an invalid value for the %q field at position %d!", unmarshalTypeError.Field, unmarshalTypeError.Offset)
			return NewFailureResponse(http.StatusBadRequest, errorMessage)

		// Handle unknown fields.
		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			errorMessage := fmt.Sprintf("Request body contains unknown field '%s'!", fieldName)
			return NewFailureResponse(http.StatusBadRequest, errorMessage)

		// Handle empty request body.
		case errors.Is(err, io.EOF):
			errorMessage := "Request body must not be empty!"
			return NewFailureResponse(http.StatusBadRequest, errorMessage)

		// Handle too large body.
		case err.Error() == "http: request body too large":
			errorMessage := "Request body must not be larger than 1KB!"
			return NewFailureResponse(http.StatusRequestEntityTooLarge, errorMessage)

		}
	}
	defer r.Body.Close()

	// Handle if client tries to send more than one JSON object.
	if err := decoder.Decode(&struct{}{}); err != io.EOF {
		errorMessage := "Request body must only contain a single JSON object!"
		return NewFailureResponse(http.StatusBadRequest, errorMessage)
	}

	// Validate input.
	if err := validator.New().Struct(dst); err != nil {
		return NewFailureResponse(http.StatusBadRequest, err.Error())
	}

	// If everything goes well, don't return anything.
	return nil
}
