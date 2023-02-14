package stripehandlers

import (
	"context"
	"encoding/json"
	"fmt"

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
			CheckoutSessionCompletes: events.handleCheckoutSessionComplete,
			ChargeSuccedded:          events.handleChargeSuccedded,
			PaymentIntentSuccedded:   events.handlePaymentIntentSuccedded,
			PaymentIntentCreated:     events.handlePaymentIntentCreated,
		},
	}
}

// const (
// errNoMessageFound = "no message found in payment intent"
// )

// handlePaymentIntentSuccedded handles stripe payment_intent.succeeded event
func (e *EventHandlers) handlePaymentIntentSuccedded(ctx context.Context, data stripe.EventData) error {
	var paymentIntent stripe.PaymentIntent
	err := json.Unmarshal(data.Raw, &paymentIntent)
	if err != nil {
		return fmt.Errorf("error unmarshaling event data: %w", err)
	}

	// pj, _ := json.MarshalIndent(paymentIntent, "", "    ")
	// fmt.Printf("%+v\n", string(pj))

	fmt.Printf("META: %+v\n", paymentIntent.Metadata)

	// tx := models.Transaction{
	// 	Amount:  float64(paymentIntent.Amount),
	// 	Message: message,
	// }
	//
	// result := e.db.WithContext(ctx).Create(&tx)
	// if result.Error != nil {
	// 	return fmt.Errorf("error unable to commit transition to db: %w", result.Error)
	// }

	return nil
}

// handleCheckoutSessionComplete handles stripe checkout.session.complete event
func (e *EventHandlers) handleCheckoutSessionComplete(ctx context.Context, data stripe.EventData) error {
	var checkoutSession stripe.CheckoutSession
	err := json.Unmarshal(data.Raw, &checkoutSession)
	if err != nil {
		return fmt.Errorf("error unmarshaling event data: %w", err)
	}

	// pj, _ := json.MarshalIndent(checkoutSession, "", "    ")
	// fmt.Printf("%+v\n", string(pj))

	fmt.Printf("META: %+v\n", checkoutSession.Metadata)

	// message, ok := checkoutSession.Metadata["message"]
	// if !ok {
	// 	return fmt.Errorf(errNoMessageFound)
	// }

	// tx := models.Transaction{
	// 	Amount:  float64(checkoutSession.Amount),
	// 	Message: message,
	// }
	//
	// result := e.db.WithContext(ctx).Create(&tx)
	// if result.Error != nil {
	// 	return fmt.Errorf("error unable to commit transition to db: %w", result.Error)
	// }

	return nil
}

// handleChargeSuccedded handles stripe charge.succeeded event
func (e *EventHandlers) handleChargeSuccedded(ctx context.Context, data stripe.EventData) error {
	var charge stripe.Charge
	err := json.Unmarshal(data.Raw, &charge)
	if err != nil {
		return fmt.Errorf("error unmarshaling event data: %w", err)
	}

	// pj, _ := json.MarshalIndent(charge, "", "    ")
	// fmt.Printf("%+v\n", string(pj))

	fmt.Printf("META: %+v\n", charge.Metadata)

	// message, ok := checkoutSession.Metadata["message"]
	// if !ok {
	// 	return fmt.Errorf(errNoMessageFound)
	// }

	// tx := models.Transaction{
	// 	Amount:  float64(checkoutSession.Amount),
	// 	Message: message,
	// }
	//
	// result := e.db.WithContext(ctx).Create(&tx)
	// if result.Error != nil {
	// 	return fmt.Errorf("error unable to commit transition to db: %w", result.Error)
	// }

	return nil
}

// handlePaymentIntentCreated handles stripe charge.succeeded event
func (e *EventHandlers) handlePaymentIntentCreated(ctx context.Context, data stripe.EventData) error {
	var paymentIntent stripe.PaymentIntent
	err := json.Unmarshal(data.Raw, &paymentIntent)
	if err != nil {
		return fmt.Errorf("error unmarshaling event data: %w", err)
	}

	// pj, _ := json.MarshalIndent(paymentIntent, "", "    ")
	// fmt.Printf("%+v\n", string(pj))

	fmt.Printf("META: %+v\n", paymentIntent.Metadata)

	// message, ok := checkoutSession.Metadata["message"]
	// if !ok {
	// 	return fmt.Errorf(errNoMessageFound)
	// }

	// tx := models.Transaction{
	// 	Amount:  float64(checkoutSession.Amount),
	// 	Message: message,
	// }
	//
	// result := e.db.WithContext(ctx).Create(&tx)
	// if result.Error != nil {
	// 	return fmt.Errorf("error unable to commit transition to db: %w", result.Error)
	// }

	return nil
}
