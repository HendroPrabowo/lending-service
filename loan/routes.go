package loan

import (
	"github.com/go-chi/chi/v5"
	"github.com/newrelic/go-agent/v3/newrelic"

	"lending-service/account"
	"lending-service/config/monitoring"
)

type routes struct {
	controller controller
	middleware account.Middleware
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
		r.Post(newrelic.WrapHandleFunc(monitoring.NewrelicApp, "/api/v1/loan", routes.controller.AddLoan))
	})
}
