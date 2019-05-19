package exception

import (
	"encoding/json"
	"net/http"
)

// Exception message
type Exception struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

// HandleError response
func HandleError(w http.ResponseWriter, e Exception) {
	w.WriteHeader(e.Code)
	json.NewEncoder(w).Encode(e)
}
