package utils

import (
	"encoding/json"
	"net/http"

	"github.com/sahilrana7582/multi-tenent-e-com-user-service/internal/dto"
)

func WriteJSON(w http.ResponseWriter, status int, msg string, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(dto.APIResponse{
		Status:  status,
		Message: msg,
		Data:    data,
	})
}
