package http

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetupRouter(r *chi.Mux, orderHandler *OrderHandler, authMiddleware func(http.Handler) http.Handler) {
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Group(func(r chi.Router) {
		r.Use(authMiddleware)
		r.Route("/orders", func(r chi.Router) {
			r.Post("/", orderHandler.CreateOrder)
			r.Get("/{id}", orderHandler.GetOrderById)
			r.Get("/user/{userId}", orderHandler.GetOrdersByUserId)
			r.Put("/{id}", orderHandler.UpdateOrderStatus)
			r.Delete("/{id}", orderHandler.DeleteOrder)
		})
	})
}
