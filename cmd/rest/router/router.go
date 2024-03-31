package router

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"

	"capital-gains-api/cmd/rest/handler"
)

type Router struct {
	port                string
	timeout             time.Duration
	router              *chi.Mux
	operationTaxHandler *handler.Handler
}

func NewHTTPRouter(operationTaxHandler *handler.Handler) *Router {
	router := &Router{
		router:              chi.NewRouter(),
		port:                "8080",
		operationTaxHandler: operationTaxHandler,
		timeout:             3000,
	}

	router.routes()

	return router
}

func (r *Router) routes() {
	r.operationTaxRouters()
}

func (r *Router) operationTaxRouters() {

	r.router.Route("/operation", func(router chi.Router) {
		router.Post("/tax", r.operationTaxHandler.PostTaxOperation)
	})

}

func (r *Router) ListenAndServe() error {
	fmt.Println("service running on 8080 http port...")

	server := &http.Server{
		Addr:              ":8080",
		ReadHeaderTimeout: 3000,
		Handler:           r.router,
	}

	return server.ListenAndServe()
}
