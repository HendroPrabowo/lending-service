package account

import (
	"github.com/go-chi/chi/v5"
	"github.com/newrelic/go-agent/v3/newrelic"

	"lending-service/config/monitoring"
)

type routes struct {
	controller controller
	middleware middleware
}

func newRoutes(controller controller, middleware middleware) routes {
	return routes{
		controller: controller,
		middleware: middleware,
	}
}

func (routes routes) RegisterRoutes(r *chi.Mux) {
	r.Post(newrelic.WrapHandleFunc(monitoring.NewrelicApp, "/api/v1/register", routes.controller.Register))
	r.Post(newrelic.WrapHandleFunc(monitoring.NewrelicApp, "/api/v1/login", routes.controller.Login))

	r.Group(func(r chi.Router) {
		r.Use(routes.middleware.validateToken)
		r.Put(newrelic.WrapHandleFunc(monitoring.NewrelicApp, "/api/v1/update", routes.controller.Update))
	})
}
