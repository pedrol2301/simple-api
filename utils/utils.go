package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func ParseJsonBody(body io.Reader, v any) error {
	decoder := json.NewDecoder(body)
	if err := decoder.Decode(v); err != nil {
		return fmt.Errorf("error while parsing json body: %v", err)
	}
	return nil
}

func WriteJsonResponse(w http.ResponseWriter, statuscode int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statuscode)

	return json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, statuscode int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statuscode)

	json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
}
