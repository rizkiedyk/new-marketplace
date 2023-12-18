package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type Order struct {
	Id             string `gorm:"size:36;not null;uniqueIndex;primary_key"`
	UserID         string `gorm:"size:36;index"`
	User           User
	OrderItems     []OrderItem
	OrderCostumer  *OrderCostumer
	Code           string `gorm:"size:50;index"`
	Status         int
	OrderDate      time.Time
	PaymentDue     time.Time
	PaymentStatus  string          `gorm:"size:50;index"`
	PaymentToken   string          `gorm:"size:50;index"`
	BaseTotalPrice decimal.Decimal `gorm:"type:decimal(16,2)"`
	TaxAmount      decimal.Decimal `gorm:"type:decimal(16,2)"`
}
