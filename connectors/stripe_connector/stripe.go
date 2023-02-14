package stripeconnector

import (
	"errors"
	"fmt"

	stripe "github.com/stripe/stripe-go/v74"
	stripeSession "github.com/stripe/stripe-go/v74/checkout/session"
)

const (
	gbpCurrency = "GBP"
)

type Product struct {
	Description string
	Name        string
	MetaData    map[string]string
	Price       float64
}

type StripeConnector struct {
	publicKey  string
	secretKey  string
	successURL string
	cancelURL  string
}

func NewStripeConnector(publicKey, secretKey, successURL, cancelURL string) *StripeConnector {
	stripe.Key = secretKey

	return &StripeConnector{
		publicKey:  publicKey,
		secretKey:  secretKey,
		successURL: successURL,
		cancelURL:  cancelURL,
	}

}

func (s *StripeConnector) GetPublicKey() string {
	return s.publicKey
}

func (s *StripeConnector) StartCheckoutSession(product Product) (*stripe.CheckoutSession, error) {

	if product.Price < 0 {
		return nil, errors.New("price cannot be negative")
	}

	if product.Price > 100 {
		return nil, errors.New("cannot charge more than 100")
	}

	priceInPennies := int64(product.Price * 100)

	params := &stripe.CheckoutSessionParams{
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
					Currency:    stripe.String(gbpCurrency),
					UnitAmount:  &priceInPennies,
					TaxBehavior: stripe.String("inclusive"),
					ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
						Description: &product.Description,
						Name:        &product.Name,
					},
				},
				Quantity: stripe.Int64(1),
			},
		},
		PaymentIntentData: &stripe.CheckoutSessionPaymentIntentDataParams{
			Metadata: product.MetaData,
		},
		Mode:         stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL:   stripe.String(s.successURL),
		CancelURL:    stripe.String(s.cancelURL),
		AutomaticTax: &stripe.CheckoutSessionAutomaticTaxParams{Enabled: stripe.Bool(true)},
	}

	ss, err := stripeSession.New(params)
	if err != nil {
		return nil, fmt.Errorf("error creating stripe session: %s", err)
	}

	return ss, nil
}
