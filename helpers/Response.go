// Package helpers implements commonly used functions (response API)//
package helpers

import (
	"encoding/json"
	"net/http"
)

// Response is
func Response(w http.ResponseWriter, httpStatus int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	json.NewEncoder(w).Encode(data)
}
