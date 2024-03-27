package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

type Router struct {
	port                string
	timeout             time.Duration
	cors                cors.Options
	router              *chi.Mux
	operationTaxHandler *Handler
}

func NewHTTPRouter(operationTaxHandler *Handler) *Router {
	router := &Router{
		router:              chi.NewRouter(),
		port:                "8080",
		operationTaxHandler: operationTaxHandler,
		timeout:             3000,
	}

	// router.corsOptions()
	// router.Middlewares()
	router.routes()

	return router
}

func (r *Router) routes() {
	r.operationTaxRouters()
}

func (r *Router) operationTaxRouters() {

	r.router.Route("operation", func(router chi.Router) {
		router.Post("/tax", r.operationTaxHandler.PostTaxOperation)
	})

}

func (r *Router) ListenAndServe() error {
	fmt.Println("service running on 8080 http port...")

	server := &http.Server{
		Addr:              "8080",
		ReadHeaderTimeout: 3000,
		Handler:           r.router,
	}

	return server.ListenAndServe()
}
