package repository

import (
	"context"

	"github.com/e-ziswaf/eziswaf-api/internal/app/model"
)

type IDonorRepository interface {
	GetDonorByID(ctx context.Context, donorID uint64) (donor model.Donor, err error)
}

type DonorRepository struct {
	opt Option
}

func NewDonorRepository(opt Option) IDonorRepository {
	return &DonorRepository{
		opt: opt,
	}
}

func (dr *DonorRepository) GetDonorByID(ctx context.Context, donorID uint64) (donor model.Donor, err error) {
	result := dr.opt.DbPostgre.Where("id = ?", donorID).First(&donor)
	if result.Error != nil {
		// TODO: log
		err = result.Error
	}

	return
}
