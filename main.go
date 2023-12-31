package main

import (
	"net/http"
	"os"

	"lending-service/account"
	"lending-service/config/database"
	"lending-service/config/monitoring"
	"lending-service/constant"
	"lending-service/files"
	"lending-service/health"
	"lending-service/loan"

	"github.com/bugsnag/bugsnag-go/v2"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-chi/httplog"
	log "github.com/sirupsen/logrus"
)

func init() {
	monitoring.InitLogger()
	monitoring.InitNewRelic()
	monitoring.InitBugsnag()
	database.InitPostgreOrm()
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := chi.NewRouter()
	r = setCors(r)

	logger := httplog.NewLogger(constant.APP_NAME, httplog.Options{
		JSON: true,
	})
	r.Use(httplog.RequestLogger(logger))

	// REGISTER ALL ROUTES HERE
	// health check routes
	health.RegisterRoutes(r)

	accountRoutes, _ := account.InitializeAccount()
	accountRoutes.RegisterRoutes(r)

	loanRoutes, _ := loan.InitializeLoan()
	loanRoutes.RegisterRoutes(r)

	fileRoutes, _ := files.InitializeRoutes()
	fileRoutes.RegisterRoutes(r)

	log.Info("running on port : " + port)
	http.ListenAndServe(":"+port, bugsnag.Handler(r))
}

func setCors(r *chi.Mux) *chi.Mux {
	// Basic CORS
	// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
	cors := cors.New(cors.Options{
		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	r.Use(cors.Handler)
	return r
}
