package stripehandlers

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/JamesTiberiusKirk/money-waste/models"
	"github.com/sirupsen/logrus"
	"github.com/stripe/stripe-go/v74"
	"gorm.io/gorm"
)

type StripeEventType string

const (
	PaymentIntentSuccedded   StripeEventType = "payment_intent.succeeded"
	CheckoutSessionCompletes StripeEventType = "checkout.session.completed"
	ChargeSuccedded          StripeEventType = "charge.succeeded"
	PaymentIntentCreated     StripeEventType = "payment_intent.created"
)

type ConfigMap struct {
	events   *EventHandlers
	handlers map[StripeEventType]func(ctx context.Context, data stripe.EventData) error
}

type EventHandlers struct {
	db *gorm.DB
}

func (c *ConfigMap) Handle(ctx context.Context, event stripe.Event) error {
	handler, ok := c.handlers[StripeEventType(event.Type)]
	if !ok || handler == nil {
		return fmt.Errorf("event %s not implemented, skipping", event.Type)
	}

	if event.Data == nil {
		return fmt.Errorf("event %s has no data", event.Type)
	}

	logrus.Printf("event type, %s", event.Type)

	err := handler(ctx, *event.Data)
	if err != nil {
		return fmt.Errorf("handler %s returned with error %w", event.Type, err)
	}

	return nil
}

func NewConfigMap(db *gorm.DB) *ConfigMap {
	events := &EventHandlers{
		db: db,
	}

	return &ConfigMap{
		events: events,
		handlers: map[StripeEventType]func(context.Context, stripe.EventData) error{
			ChargeSuccedded: events.handleChargeSuccedded,
		},
	}
}

const (
	errNoMessageFound = "no message found in payment intent"
	messageMetaTag    = "message"
)

// handleChargeSuccedded handles stripe charge.succeeded event
func (e *EventHandlers) handleChargeSuccedded(ctx context.Context, data stripe.EventData) error {
	var charge stripe.Charge
	err := json.Unmarshal(data.Raw, &charge)
	if err != nil {
		return fmt.Errorf("error unmarshaling event data: %w", err)
	}

	tx := models.Transaction{
		Amount:          float64(charge.Amount) / 100,
		Message:         charge.Metadata[messageMetaTag],
		StripeChargeID:  charge.ID,
		PaymentStatus:   models.ChargeStatus(charge.Status),
		Live:            charge.Livemode,
		TransactionTime: time.Unix(charge.Created, 0),
	}

	if charge.BillingDetails != nil {
		tx.Email = charge.BillingDetails.Email
	}

	result := e.db.WithContext(ctx).Create(&tx)
	if result.Error != nil {
		return fmt.Errorf("error unable to commit transition to db: %w", result.Error)
	}

	return nil
}
