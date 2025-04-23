package authentication

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Routes(service *Service) http.Handler {
	r := chi.NewRouter()
	handler := NewHandler(service)
	r.Get("/token", handler.TokenHandler)
	return r
}
