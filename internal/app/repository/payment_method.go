package repository

import (
	"context"

	"github.com/e-ziswaf/eziswaf-api/internal/app/model"
	"gorm.io/gorm"
)

type IPaymentMethodRepository interface {
	GetEnabledPaymentMethods(ctx context.Context) (paymentMethods []model.PaymentMethod, err error)
}

type PaymentMethodRepository struct {
	opt Option
}

func NewPaymentMethodRepository(opt Option) IPaymentMethodRepository {
	return &PaymentMethodRepository{
		opt: opt,
	}
}

func (pm *PaymentMethodRepository) GetEnabledPaymentMethods(ctx context.Context) (paymentMethods []model.PaymentMethod, err error) {
	result := pm.opt.DbPostgre.WithContext(ctx).Find(&paymentMethods).Where("is_enabled = ?", true)
	if result.Error != nil {
		err = result.Error
		return
	}

	if result.RowsAffected < 1 {
		err = gorm.ErrRecordNotFound
		return
	}

	return
}
