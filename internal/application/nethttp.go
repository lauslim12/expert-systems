package application

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"
	"github.com/lauslim12/expert-systems/pkg/inference"
)

// Application constants for application modes.
const (
	applicationModeDevelopment = "development"
	applicationModeProduction  = "production"
)

// SuccessResponse is used to handle successful responses.
type SuccessResponse struct {
	Status  string      `json:"status"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// FailureResponse is used to handle failed requests.
type FailureResponse struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// NewSuccessResponse will create an instance of 'SuccessResponse' with default values.
func NewSuccessResponse(code int, message string, data interface{}) *SuccessResponse {
	return &SuccessResponse{
		Status:  "success",
		Code:    code,
		Message: message,
		Data:    data,
	}
}

// NewFailureResponse will create an instance of 'FailureResponse' with default values.
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

// Initialize application.
// The 'pathToWebDirectory' is usually filled with './web/build' (the built React application location).
func Configure(pathToWebDirectory, applicationMode string) http.Handler {
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

	// Set up custom middleware. Don't forget to inject dependencies.
	r.Use(customHeaders)
	r.Use(httpsRedirect(applicationMode))

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
			input := &inference.Input{}

			// Check input length.
			r.Body = http.MaxBytesReader(w, r.Body, 10240)

			// Create JSON decoder.
			decoder := json.NewDecoder(r.Body)
			decoder.DisallowUnknownFields()

			// Parse JSON.
			if err := decoder.Decode(input); err != nil {
				sendFailureResponse(w, NewFailureResponse(http.StatusBadRequest, err.Error()))
				return
			}

			// Perform inference with our Expert System based on the given input.
			inferredData := inference.Infer(input)

			// Send back response.
			res := NewSuccessResponse(http.StatusOK, "Successfully processed data in the Expert System!", inferredData)
			sendSuccessResponse(w, res)
		})

		// Declare method not allowed as fallback route.
		r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
			errorMessage := fmt.Sprintf("Method '%s' is not allowed in this route!", r.Method)
			sendFailureResponse(w, NewFailureResponse(http.StatusMethodNotAllowed, errorMessage))
		})

		// For this, we declare a 404.
		r.NotFound(func(w http.ResponseWriter, r *http.Request) {
			errorMessage := fmt.Sprintf("Route '%s' does not exist in this server!", r.RequestURI)
			sendFailureResponse(w, NewFailureResponse(http.StatusNotFound, errorMessage))
		})
	})

	// Fallback route, serve React app.
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

	// Returns our router instance.
	return r
}
