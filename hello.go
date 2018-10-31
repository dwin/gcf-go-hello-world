package helloworld

import (
	"encoding/json"
	"log"
	"net/http"
)

// HelloWorld ...
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// JSON Encoder
	enc := json.NewEncoder(w)
	// Create JSON Source
	b := map[string]interface{}{
		"Message": "Hello World!",
	}
	// Encode JSON Body
	if err := enc.Encode(&b); err != nil {
		log.Printf("JSON Encode error %s\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	return
}
