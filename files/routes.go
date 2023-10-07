package files

import (
	"github.com/go-chi/chi/v5"

	"lending-service/account"
)

type routes struct {
	middleware account.Middleware
	controller controller
}

func newRoutes(controller controller, middleware account.Middleware) routes {
	return routes{
		controller: controller,
		middleware: middleware,
	}
}

func (routes routes) RegisterRoutes(r *chi.Mux) {
	r.Group(func(r chi.Router) {
		r.Use(routes.middleware.ValidateToken)
		r.Get("/api/v1/file", routes.controller.Download)
		r.Post("/api/v1/file", routes.controller.Upload)
	})
}
