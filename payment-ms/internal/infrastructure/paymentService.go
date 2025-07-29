package infrastructure

import (
	"github.com/stripe/stripe-go/v78"
	"github.com/stripe/stripe-go/v78/paymentlink"
	"github.com/stripe/stripe-go/v78/price"
	"github.com/stripe/stripe-go/v78/product"
)

type PaymentService struct{}

func NewPaymentService(secretKey string) *PaymentService {
	stripe.Key = secretKey
	return &PaymentService{}
}

func (s *PaymentService) CreatePaymentLink(orderID string, amount float64, currency string) (string, error) {
	// 1. Create Product
	prod, err := product.New(&stripe.ProductParams{
		Name: stripe.String("Order " + orderID),
	})
	if err != nil {
		return "", err
	}

	// 2. Create Price
	prc, err := price.New(&stripe.PriceParams{
		Product:    stripe.String(prod.ID),
		UnitAmount: stripe.Int64(int64(amount * 100)),
		Currency:   stripe.String(currency),
	})
	if err != nil {
		return "", err
	}

	// 3. Create Payment Link with metadata
	link, err := paymentlink.New(&stripe.PaymentLinkParams{
		LineItems: []*stripe.PaymentLinkLineItemParams{
			{
				Price:    stripe.String(prc.ID),
				Quantity: stripe.Int64(1),
			},
		},
		Metadata: map[string]string{
			"order_id": orderID,
		},
	})
	if err != nil {
		return "", err
	}

	return link.URL, nil
}
