package health

import (
	"github.com/go-chi/chi/v5"
	"github.com/hellofresh/health-go/v5"
	"github.com/newrelic/go-agent/v3/newrelic"

	"lending-service/config/monitoring"
	"lending-service/constant"
)

func RegisterRoutes(r *chi.Mux) {
	h, _ := health.New(health.WithComponent(health.Component{
		Name:    constant.APP_NAME,
		Version: "v1.0",
	}))

	r.Get(newrelic.WrapHandleFunc(monitoring.NewrelicApp, "/status", h.HandlerFunc))
}
