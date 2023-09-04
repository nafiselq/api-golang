package model

import (
	"database/sql"
	"time"
)

type Campaign struct {
	ID                  uint64
	LembagaID           uint64
	StatusID            uint64
	Title               string
	TotalDonationAmount float64
	Description         sql.NullString
	BannerURL           sql.NullString
	CreatedAt           time.Time
	UpdatedAt           time.Time
	LastUpdated         time.Time
}

func (Campaign) TableName() string {
	return "campaign"
}
