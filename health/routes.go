package health

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/hellofresh/health-go/v5"
	"github.com/newrelic/go-agent/v3/newrelic"

	"lending-service/config/monitoring"
	"lending-service/constant"
	"lending-service/utility/response"
	"lending-service/utility/wraped_error"
)

func RegisterRoutes(r *chi.Mux) {
	h, _ := health.New(health.WithComponent(health.Component{
		Name:    constant.APP_NAME,
		Version: "v1.0",
	}))

	r.Get(newrelic.WrapHandleFunc(monitoring.NewrelicApp, "/", func(w http.ResponseWriter, r *http.Request) {
		resp := map[string]interface{}{
			"app_name": "welcome lending-service",
		}
		response.Ok(w, resp)
	}))
	r.Get(newrelic.WrapHandleFunc(monitoring.NewrelicApp, "/status", h.HandlerFunc))
	r.Get(newrelic.WrapHandleFunc(monitoring.NewrelicApp, "/error", func(w http.ResponseWriter, r *http.Request) {
		err := wraped_error.WrapError(fmt.Errorf("testing error"), http.StatusInternalServerError)
		response.ErrorWrapped(w, err)
	}))
}
