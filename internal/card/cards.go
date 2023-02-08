package card

import (
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/paymentintent"
)

type Card struct {
	// holds ref to Stripe sk_ value
	Secret string

	// holds ref to Stripe pk_ value
	Key      string
	Currency string
}

type Transaction struct {
	// Id
	// holds transaction status id
	Id int

	// Amount
	// represents transaction amount, avoids float point
	Amount int

	// Currency
	// holds info on transaction currency in USD like format
	Currency string

	// Last4
	// holds card last four digits
	Last4 string

	// ReturnCode
	// holds transaction bank return code from Stripe API
	ReturnCode string
}

// CreatePaymentIntent performs charge logic
// may return error code message from Stripe API
func (c *Card) CreatePaymentIntent(currency string, amount int) (*stripe.PaymentIntent, string, error) {
	// msg holds error code if any
	// specified in stripe api
	// https://stripe.com/docs/error-codes
	var msg string

	stripe.Key = c.Secret

	// create payment intent
	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(int64(amount)),
		Currency: stripe.String(currency),
	}

	pi, err := paymentintent.New(params)
	if err != nil {
		if stripeErr, ok := err.(*stripe.Error); ok {
			msg = string(stripeErr.Code)

			return nil, msg, err
		}
	}

	return pi, msg, nil
}
