package repository

import (
	"github.com/e-ziswaf/eziswaf-api/internal/app/commons"
	"gorm.io/gorm"
)

// Option anything any repo object needed
type Option struct {
	commons.Options
	DB *gorm.DB
}

// Repository all repo object injected here
type Repository struct {
	Cache  ICacheRepository
	Person IPersonRepository
	Donor  IDonorRepository
}
