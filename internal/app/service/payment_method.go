package service

import (
	"context"

	"github.com/e-ziswaf/eziswaf-api/internal/app/payloads"
)

type IPaymentMethodService interface {
	GetEnabledPaymentMethodList(ctx context.Context) (resp []payloads.PaymentMethodResp, err error)
}

type PaymentMethodService struct {
	opt Option
}

func NewPaymentMethodService(opt Option) IPaymentMethodService {
	return &PaymentMethodService{
		opt: opt,
	}
}

func (pm *PaymentMethodService) GetEnabledPaymentMethodList(ctx context.Context) (resp []payloads.PaymentMethodResp, err error) {
	paymentMethods, err := pm.opt.PaymentMethod.GetEnabledPaymentMethods(ctx)
	if err != nil {
		// TODO: add log
		return
	}

	resp = make([]payloads.PaymentMethodResp, len(paymentMethods))
	for i, paymentMethod := range paymentMethods {
		resp[i].ID = paymentMethod.ID
		resp[i].Name = paymentMethod.Name
		resp[i].BankCode = paymentMethod.BankCode
		resp[i].Type = paymentMethod.BankType
		resp[i].IsEnabled = paymentMethod.IsEnabled
		resp[i].FixedFee = paymentMethod.FixedFee
		resp[i].VariableFee = paymentMethod.VariableFee
		resp[i].Logo = paymentMethod.Logo
	}

	return
}
