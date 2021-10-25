package application

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"
	"github.com/go-playground/validator/v10"
	"github.com/lauslim12/asuka"
	"github.com/lauslim12/expert-systems/pkg/inference"
)

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
			statusCode, err := asuka.Parse(w, r, input, asuka.DefaultConfig())
			if err != nil {
				sendFailureResponse(w, NewFailureResponse(statusCode, err.Error()))
				return
			}

			// Validate fields.
			if err := validator.New().Struct(input); err != nil {
				sendFailureResponse(w, NewFailureResponse(http.StatusBadRequest, err.Error()))
				return
			}

			// Perform inference with our Expert System.
			exampleInferredData, err := inference.Infer()
			if err != nil {
				sendFailureResponse(w, NewFailureResponse(http.StatusInternalServerError, err.Error()))
				return
			}

			// Send back response.
			res := NewSuccessResponse(http.StatusOK, "Successfully processed data in the Expert System!", exampleInferredData)
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

	// Returns our router instance.
	return r
}
