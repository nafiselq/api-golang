package repository

import (
	"context"

	"github.com/e-ziswaf/eziswaf-api/internal/app/model"
)

type ILembagaRepository interface {
	GetLembagaByID(ctx context.Context, lembagaID uint64) (lembaga model.Lembaga, err error)
}

type LembagaRepository struct {
	opt Option
}

func NewLembagaRepository(opt Option) ILembagaRepository {
	return &LembagaRepository{
		opt: opt,
	}
}

func (lr *LembagaRepository) GetLembagaByID(ctx context.Context, lembagaID uint64) (lembaga model.Lembaga, err error) {
	result := lr.opt.DbPostgre.WithContext(ctx).First(&lembaga, lembagaID)
	if result.Error != nil {
		err = result.Error
		return
	}

	return
}
