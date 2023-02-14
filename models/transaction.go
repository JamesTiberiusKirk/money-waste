package models

import "gorm.io/gorm"

type ChargeStatus string

const (
	ChargeStatusFailed    ChargeStatus = "failed"
	ChargeStatusPending   ChargeStatus = "pending"
	ChargeStatusSucceeded ChargeStatus = "succeeded"
)

type Transaction struct {
	gorm.Model
	StripeChargeID string `gorm:"uniqueIndex"`
	PaymentStatus  ChargeStatus
	Message        string
	Amount         float64
	Email          string
	Live           bool
}
