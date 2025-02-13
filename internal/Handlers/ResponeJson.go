package handlers

import (
	"encoding/json"
	"net/http"
	"real-time-forum/internal/models/entities"
	"time"
)

func SendResponse(w http.ResponseWriter, task, message string, isSuccess bool, statusCode int,data ...[]entities.Post) {
	response := map[string]interface{}{
		"Task":    task,
		"Success": isSuccess,
		"message": message,
        "Time":    time.Now().Format(time.RFC3339),
	}
    if len(data) > 0 {
        response["Posts"] = data // Directly assign the slice
    }

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization")
}
