package route

import (
	"net/http"

	stripehandlers "github.com/JamesTiberiusKirk/money-waste/events/stripe_handlers"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/stripe/stripe-go/v74"
)

const (
	stripeWebhookAPIRoute = "/stripe/stripe_webhooks"
)

// UserRoute user route dependency struct.
type StripeWebhookRoute struct {
	events *stripehandlers.ConfigMap
}

// NewStripeWebhookRoute struct instance.
func NewStripeWebhookRoute(stripeEventHandler *stripehandlers.ConfigMap) *Route {
	depts := &StripeWebhookRoute{
		events: stripeEventHandler,
	}

	return &Route{
		RouteID:     "stripeWebhook",
		Path:        stripeWebhookAPIRoute,
		Depts:       depts,
		PostHandler: depts.PostHandler,
	}
}

func (r StripeWebhookRoute) PostHandler(c echo.Context) error {
	event := stripe.Event{}
	err := c.Bind(&event)

	if err != nil {
		logrus.
			WithError(err).
			Print("Failed to parse webhook body json")
		return c.NoContent(http.StatusBadRequest)
	}

	err = r.events.Handle(c.Request().Context(), event)
	if err != nil {
		logrus.
			WithError(err).
			Print("Failed to process webhook event")
		return c.NoContent(http.StatusBadRequest)

	}

	return c.NoContent(http.StatusOK)
}
