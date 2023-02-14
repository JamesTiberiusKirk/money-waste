package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	Message string
	Amount  float64
}
