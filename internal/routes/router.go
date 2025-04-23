package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"pustobaseproject/internal/authentication"
	"pustobaseproject/internal/domain/players"
	playermodule "pustobaseproject/internal/modules/players"
)

func SetupRouter(playerService *players.Service) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Создаём экземпляр authentication.Service
	authService := authentication.NewService()

	r.Route("/api/v1", func(api chi.Router) {
		api.Mount("/players", playermodule.Routes(playerService))
		api.Mount("/auth", authentication.Routes(authService)) // Передаём authService
	})
	return r
}
