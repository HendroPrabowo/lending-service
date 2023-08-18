package health

import (
	"github.com/go-chi/chi/v5"
	"github.com/hellofresh/health-go/v5"

	"lending-service/constant"
)


func RegisterRoutes(r *chi.Mux) {
	h, _ := health.New(health.WithComponent(health.Component{
		Name:    constant.APP_NAME,
		Version: "v1.0",
	}))

	r.Get("/status", h.HandlerFunc)
}
