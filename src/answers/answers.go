package answers

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSON writes a response in JSON to the given response writer
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			log.Fatal(err)
		}
	}

}

// Error transforms the given error into a JSON
func Error(w http.ResponseWriter, statusCode int, err error) {
	JSON(w, statusCode, struct {
		Err string `json:"error"`
	}{
		Err: err.Error(),
	})

}
