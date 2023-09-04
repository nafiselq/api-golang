package model

import "time"

type Donation struct {
	ID          uint64
	CampaignID  uint64
	StatusID    uint64
	Amount      float64
	ChargeFee   float32
	PlatformFee float32
	IsAnonymous bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	LastUpdated time.Time
}

type DonationDetail struct {
	ID                  uint64  `gorm:"column:id"`
	Amount              float64 `gorm:"column:amount"`
	StatusID            uint64  `gorm:"column:status_id"`
	StatusName          string  `gorm:"column:donation_status_name"`
	PaymentMethodID     uint64  `gorm:"column:payment_method_id"`
	PaymentMethodName   string  `gorm:"column:payment_method_name"`
	CampaignID          uint64  `gorm:"column:campaign_id"`
	CampaignTitle       string  `gorm:"column:campaign_title"`
	CampaignBannerURL   string  `gorm:"column:campaign_banner_url"`
	CampaignTotalAmount float64 `gorm:"column:campaign_total_amount"`
	CampaignLembagaID   uint64  `gorm:"column:campaign_lembaga_id"`
}
