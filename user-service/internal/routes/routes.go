package routes

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/sahilrana7582/multi-tenent-e-com-user-service/internal/handler"
)

func NewRouter(userHandler *handler.UserHandler) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Timeout(30 * time.Second))

	r.Route("/api/v1", func(r chi.Router) {

		r.Route("/users", func(r chi.Router) {
			r.Post("/register", customHandler(userHandler.RegisterUser))
			// r.Post("/login", customHandler(userHandler.LoginUser)) // add later
		})

	})

	return r
}
