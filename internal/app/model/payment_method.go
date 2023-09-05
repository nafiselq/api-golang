package model

import "time"

type PaymentMethod struct {
	ID          uint64
	Name        string
	BankCode    string
	BankType    string `gorm:"column:type"`
	IsEnabled   bool
	FixedFee    float32
	VariableFee float32
	Logo        string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (PaymentMethod) TableName() string {
	return "payment_method"
}
