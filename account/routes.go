package account

import "github.com/go-chi/chi/v5"

type routes struct {
	controller controller
}

func newRoutes(controller controller) routes {
	return routes{controller: controller}
}

func (routes routes) RegisterRoutes(r *chi.Mux) {
	r.Post("/api/v1/register", routes.controller.Register)
}