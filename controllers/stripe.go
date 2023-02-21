package controllers

import (
	"os"

	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/checkout/session"
)

func Stripe(total uint) string {
	stripe.Key = os.Getenv("Stripe")

	amount := int64(total)
	params := &stripe.CheckoutSessionParams{

		Mode: stripe.String(string(stripe.CheckoutSessionModePayment)),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			&stripe.CheckoutSessionLineItemParams{
				PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
					Currency: stripe.String("inr"),
					ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
						Name: stripe.String("products"),
					},
					UnitAmount: stripe.Int64(amount * 100), //stripe.Int64(*100),
				},
				Quantity: stripe.Int64(1),
			},
		},
		SuccessURL: stripe.String("http://localhost:8080/success"),
		CancelURL:  stripe.String("http://localhost:8080/cancel"),
	}

	s, _ := session.New(params)

	// if err != nil {
	// 	log.Fatal("session.New: %v", err)
	// }
	return s.URL

}
