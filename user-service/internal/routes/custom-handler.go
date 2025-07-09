package routes

import (
	"errors"
	"net/http"
)

func customHandler(f func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := f(w, r)
		if err != nil {
			status := http.StatusInternalServerError
			var statusErr interface {
				error
				HTTPStatus() int
			}
			if errors.As(err, &statusErr) {
				status = statusErr.HTTPStatus()
			}
			http.Error(w, err.Error(), status)
		}
	}
}
