package routes

import (
	"encoding/json"
	"errors"
	"net/http"
)

func customHandler(f func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := f(w, r)
		if err == nil {
			return
		}

		status := http.StatusInternalServerError
		message := "Internal server error"

		type httpError interface {
			error
			HTTPStatus() int
		}

		var herr httpError
		if errors.As(err, &herr) {
			status = herr.HTTPStatus()
			message = err.Error()
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  status,
			"message": message,
		})
	}
}
