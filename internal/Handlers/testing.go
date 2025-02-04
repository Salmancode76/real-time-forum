package handlers

import (
	"encoding/json"
	"net/http"
)

func S_test(w http.ResponseWriter, r *http.Request) {
	// Check if the Accept header is set to application/json
	if r.Header.Get("Accept") == "application/json" {
		// If the Accept header is correct, return JSON data
		response := map[string]interface{}{
			"showRequestMade": true,
			"message":         "Yes, you can show content",
			"routeName":       "/s",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	} else {
		// Serve HTML for the /s route
		http.ServeFile(w, r, "./index.html")
	}
}
