package http

import (
	"encoding/json"
	"io"
	"net/http"
	"payment-ms/internal/domain"
	"payment-ms/internal/dto"
	"payment-ms/pkg"

	"github.com/stripe/stripe-go/v78"
)

type PaymentHandler struct {
	usecase domain.PaymentUseCase
}

func NewPaymentHandler(uc domain.PaymentUseCase) *PaymentHandler {
	return &PaymentHandler{usecase: uc}
}

// @Summary Initiate a payment
// @Description Initiate a payment for an order
// @Accept json
// @Produce json
// @Header 200 {string} Authorization "Bearer token"
// @Param request body dto.InitiatePaymentRequest true "Payment request"
// @Security BearerAuth
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /initiate-payment [post]
func (h *PaymentHandler) InitiatePayment(w http.ResponseWriter, r *http.Request) {
	var req dto.InitiatePaymentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		WriteError(w, domain.BadRequestErr("Invalid request body"))
		return
	}

	if err := pkg.ValidateRequest(req); err != nil {
		WriteError(w, err)
		return
	}

	paymentLink, err := h.usecase.StartPayment(req.OrderID, req.Amount, req.Currency)
	if err != nil {
		WriteError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"payment_link": paymentLink})
}

// @Summary Handle a webhook
// @Description Handle a webhook from Stripe
// @Accept json
// @Produce json
// @Header 200 {string} Authorization "Bearer token"
// @Param event body map[string]interface{} true "Stripe event"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /webhook [post]
func (h *PaymentHandler) HandleWebhook(w http.ResponseWriter, r *http.Request) {
	payload, err := io.ReadAll(r.Body)
	if err != nil {
		WriteError(w, domain.BadRequestErr("Invalid request body"))
		return
	}

	var event stripe.Event
	err = json.Unmarshal(payload, &event)
	if err != nil {
		WriteError(w, domain.BadRequestErr("Invalid request body"))
		return
	}

	switch event.Type {
	case "payment_link.completed":
		var pl stripe.PaymentLink
		err = json.Unmarshal(event.Data.Raw, &pl)
		orderID := pl.Metadata["order_id"]
		if err != nil {
			WriteError(w, domain.BadRequestErr("Invalid request body"))
			return
		}

		h.usecase.UpdatePaymentStatus(orderID, domain.PaymentStatus(event.Data.Object["status"].(string)))
		w.WriteHeader(http.StatusOK)
		return
	case "payment_link.expired":
		var pl stripe.PaymentLink
		err = json.Unmarshal(event.Data.Raw, &pl)
		orderID := pl.Metadata["order_id"]
		if err != nil {
			WriteError(w, domain.BadRequestErr("Invalid request body"))
			return
		}

		h.usecase.UpdatePaymentStatus(orderID, domain.PaymentStatus(event.Data.Object["status"].(string)))
		w.WriteHeader(http.StatusOK)
		return

	case "payment_link.canceled":
		var pl stripe.PaymentLink
		err = json.Unmarshal(event.Data.Raw, &pl)
		orderID := pl.Metadata["order_id"]
		if err != nil {
			WriteError(w, domain.BadRequestErr("Invalid request body"))
			return
		}

		h.usecase.UpdatePaymentStatus(orderID, domain.PaymentStatus(event.Data.Object["status"].(string)))
		w.WriteHeader(http.StatusOK)
		return
	}

	WriteError(w, domain.BadRequestErr("Invalid event type"))
}

func WriteError(w http.ResponseWriter, err error) {
	if appErr, ok := err.(*domain.AppError); ok {
		http.Error(w, appErr.Message, appErr.StatusCode)
	} else {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}
