package application

import (
	"encoding/json"
	"net/http"
)

// Application constants for modes.
const (
	applicationModeDevelopment = "development"
	applicationModeProduction  = "production"
)

// SuccessResponse is used to handle successful responses.
type SuccessResponse struct {
	Status  string      `json:"status"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
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
