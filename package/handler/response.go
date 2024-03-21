package handler

import (
	"encoding/json"
	"net/http"
	logger "testTask_retail/logs"
)

type Err struct {
	Message string `json:"message"`
}

type StatusResponse struct {
	Status string `json:"status"`
}

func NewErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	logger.Log.Error(message)
	w.WriteHeader(statusCode)
	response := map[string]interface{}{
		"error": "message",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
