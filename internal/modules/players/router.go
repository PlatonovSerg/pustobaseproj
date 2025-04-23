package playermodule

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"pustobaseproject/internal/domain/players"
	"pustobaseproject/internal/middleware"
)

func Routes(service *players.Service) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.PlayerMiddleware(service)) // подключаем middleware

	r.Get("/save", GetPlayerHandler) // GET /api/v1/players/me
	return r
}
