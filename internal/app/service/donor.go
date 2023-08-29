package service

import (
	"context"
	"errors"

	"github.com/e-ziswaf/eziswaf-api/internal/app/model"
)

type IDonorService interface {
	GetDonorByID(ctx context.Context, donorID uint64) (donor model.Donor, err error)
}

type DonorService struct {
	opt Option
}

func NewDonorService(opt Option) IDonorService {
	return &DonorService{
		opt: opt,
	}
}

func (ds *DonorService) GetDonorByID(ctx context.Context, donorID uint64) (donor model.Donor, err error) {
	donor, err = ds.opt.Repository.Donor.GetDonorByID(ctx, donorID)
	if err != nil {
		// TODO: manipulate error for user facing
		err = errors.New("gagal mendapatkan data donor. silakan coba lagi")
	}

	return
}
