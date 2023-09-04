package model

import "time"

type Lembaga struct {
	ID        uint64
	StatusID  uint64
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Lembaga) TableName() string {
	return "lembaga"
}
