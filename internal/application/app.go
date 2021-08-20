package application

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/httprate"
	"github.com/go-playground/validator"
)

// Person is an object that represents the request body.
type Person struct {
	Name    string `json:"name" validate:"required"`
	Address string `json:"address" validate:"required"`
}

// SuccessResponse is used to handle successful responses.
type SuccessResponse struct {
	Status  string  `json:"status"`
	Code    int     `json:"code"`
	Message string  `json:"message"`
	Data    *Person `json:"data,omitempty"`
}

func NewSuccessResponse(code int, message string, data *Person) *SuccessResponse {
	return &SuccessResponse{
		Status:  "success",
		Code:    code,
		Message: message,
		Data:    data,
	}
}

// FailureResponse is used to handle failed requests.
type FailureResponse struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewFailureResponse(code int, message string) *FailureResponse {
	return &FailureResponse{
		Status:  "fail",
		Code:    code,
		Message: message,
	}
}

// Utility function to send back success response to users.
func sendSuccessResponse(w http.ResponseWriter, successResponse *SuccessResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(successResponse.Code)
	json.NewEncoder(w).Encode(successResponse)
}

// Utility function to send back error response to users.
func sendFailureResponse(w http.ResponseWriter, failureResponse *FailureResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(failureResponse.Code)
	json.NewEncoder(w).Encode(failureResponse)
}

// Utility function to decode a JSON request body.
func decodeJSONBody(w http.ResponseWriter, r *http.Request) (*Person, *FailureResponse) {
	// Initialize our variable to be returned - functional style.
	parsedBody := &Person{}

	// Check if Header is 'Content-Type: application/json'.
	if r.Header.Get("Content-Type") != "application/json" {
		return nil, NewFailureResponse(http.StatusUnsupportedMediaType, "The 'Content-Type' header is not 'application/json'!")
	}

	// Parse body, and set max bytes reader (1KB).
	r.Body = http.MaxBytesReader(w, r.Body, 1024)
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(parsedBody); err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError

		switch {
		// Handle syntax errors.
		case errors.As(err, &syntaxError):
			errorMessage := fmt.Sprintf("Request body contains a badly formatted JSON at position %d!", syntaxError.Offset)
			return nil, NewFailureResponse(http.StatusBadRequest, errorMessage)

		// Handle unexpected EOFs.
		case errors.Is(err, io.ErrUnexpectedEOF):
			errorMessage := "Request body contains a badly-formed JSON!"
			return nil, NewFailureResponse(http.StatusBadRequest, errorMessage)

		// Handle wrong data-type in request body.
		case errors.As(err, &unmarshalTypeError):
			errorMessage := fmt.Sprintf("Request body contains an invalid value for the %q field at position %d!", unmarshalTypeError.Field, unmarshalTypeError.Offset)
			return nil, NewFailureResponse(http.StatusBadRequest, errorMessage)

		// Handle unknown fields.
		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			errorMessage := fmt.Sprintf("Request body contains unknown field '%s'!", fieldName)
			return nil, NewFailureResponse(http.StatusBadRequest, errorMessage)

		// Handle empty request body.
		case errors.Is(err, io.EOF):
			errorMessage := "Request body must not be empty!"
			return nil, NewFailureResponse(http.StatusBadRequest, errorMessage)

		// Handle too large body.
		case err.Error() == "http: request body too large":
			errorMessage := "Request body must not be larger than 1KB!"
			return nil, NewFailureResponse(http.StatusRequestEntityTooLarge, errorMessage)

		// Handle other errors.
		default:
			return nil, NewFailureResponse(http.StatusInternalServerError, err.Error())
		}
	}
	defer r.Body.Close()

	// Handle if client tries to send more than one JSON object.
	if err := decoder.Decode(&struct{}{}); err != io.EOF {
		errorMessage := "Request body must only contain a single JSON object!"
		return nil, NewFailureResponse(http.StatusBadRequest, errorMessage)
	}

	// Validate input.
	if err := validator.New().Struct(parsedBody); err != nil {
		return nil, NewFailureResponse(http.StatusBadRequest, err.Error())
	}

	// If everything goes well, don't return anything.
	return parsedBody, nil
}

// Initialize application.
// The 'pathToWebDirectory' is usually filled with './web/build' (the built React application location).
func Configure(pathToWebDirectory string) http.Handler {
	// Create a Chi instance.
	r := chi.NewRouter()

	// Set up middlewares.
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set up Chi's third party libraries.
	r.Use(httprate.Limit(
		200,
		1*time.Minute,
		httprate.WithLimitHandler(func(w http.ResponseWriter, r *http.Request) {
			res := NewFailureResponse(http.StatusTooManyRequests, "You have performed too many requests! Please try again in a minute!")
			sendFailureResponse(w, res)
		}),
	))

	// Set up custom middleware.
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("X-Expert-Systems", "Miyuki")
			w.Header().Add("Server", "net/http")
			next.ServeHTTP(w, r)
		})
	})

	// Group routes.
	r.Route("/api/v1", func(r chi.Router) {
		// Use compression on API.
		r.Use(middleware.Compress(5))

		// Sample GET request.
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			res := NewSuccessResponse(http.StatusOK, "Welcome to 'net/http' API!", nil)
			sendSuccessResponse(w, res)
		})

		// Sample POST request.
		r.Post("/", func(w http.ResponseWriter, r *http.Request) {
			processedData, failureResponse := decodeJSONBody(w, r)
			if failureResponse != nil {
				sendFailureResponse(w, failureResponse)
				return
			}

			// Do Expert System function here.
			// End of Expert System function.

			// Send back response.
			res := NewSuccessResponse(http.StatusOK, "Successfully processed data in the Expert Systems!", processedData)
			sendSuccessResponse(w, res)
		})

		// Declare method not allowed as fallback route.
		r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
			errorMessage := fmt.Sprintf("Method '%s' is not allowed in this route!", r.Method)
			res := NewFailureResponse(http.StatusMethodNotAllowed, errorMessage)
			sendFailureResponse(w, res)
		})

		// For this, we declare a 404.
		r.NotFound(func(w http.ResponseWriter, r *http.Request) {
			errorMessage := fmt.Sprintf("Route '%s' does not exist in this server!", r.RequestURI)
			res := NewFailureResponse(http.StatusNotFound, errorMessage)
			sendFailureResponse(w, res)
		})
	})

	// Fallback route, serve React app. Below works, but the tests return 404.
	r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		fs := http.FileServer(http.Dir(pathToWebDirectory))

		// If there is no route in the React application, send back 404.
		// Else, send back React application.
		if _, err := os.Stat(pathToWebDirectory + r.RequestURI); os.IsNotExist(err) {
			w.WriteHeader(http.StatusNotFound)
			http.StripPrefix(r.RequestURI, fs).ServeHTTP(w, r)
		} else {
			fs.ServeHTTP(w, r)
		}
	})

	return r
}
