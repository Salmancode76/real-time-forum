package handlers

import (
	"encoding/json"
	"net/http"
	"time"
)

func SendResponse(w http.ResponseWriter, task, message string, isSuccess bool, statusCode int) {
	response := map[string]interface{}{
		"Task":    task,
		"Success": isSuccess,
		"message": message,
		"Time":    time.Now(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization")
}
