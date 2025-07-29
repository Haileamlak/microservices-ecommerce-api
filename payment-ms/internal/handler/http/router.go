package http

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetupRouter(r *chi.Mux, paymentHandler *PaymentHandler, authMiddleware func(http.Handler) http.Handler) {
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Group(func(r chi.Router) {
		r.Use(authMiddleware)
		r.Route("/initiate-payment", func(r chi.Router) {
			r.Post("/", paymentHandler.InitiatePayment)
		})
	})
	r.Post("/webhook", paymentHandler.HandleWebhook)
}
