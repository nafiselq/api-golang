package service

import (
	"context"

	"github.com/e-ziswaf/eziswaf-api/internal/app/model"
	"github.com/e-ziswaf/eziswaf-api/internal/app/payloads"
)

type IDonationService interface {
	GetDonationDetail(ctx context.Context, donationID uint64) (resp payloads.GetDonationDetailByIDResp, err error)
}

type DonationService struct {
	opt Option
}

func NewDonationService(opt Option) IDonationService {
	return &DonationService{
		opt: opt,
	}
}

func (ds *DonationService) GetDonationDetail(ctx context.Context, donationID uint64) (resp payloads.GetDonationDetailByIDResp, err error) {
	donationDetail, err := ds.opt.Repository.Donation.GetDonationDetailByID(ctx, donationID)
	if err != nil {
		// TODO: add lo
		return
	}

	lembaga, err := ds.opt.Repository.Lembaga.GetLembagaByID(ctx, donationDetail.CampaignLembagaID)
	if err != nil {
		// TODO: add log
	}

	resp = buildDonationDetailResponse(donationDetail, lembaga)

	return
}

func buildDonationDetailResponse(donation model.DonationDetail, lembaga model.Lembaga) (result payloads.GetDonationDetailByIDResp) {
	result.ID = donation.ID
	result.Amount = donation.Amount
	result.Status.ID = donation.StatusID
	result.Status.Name = donation.StatusName
	result.PaymentMethod.ID = donation.PaymentMethodID
	result.PaymentMethod.Name = donation.PaymentMethodName
	result.CampaignResp.ID = donation.CampaignID
	result.CampaignResp.Title = donation.CampaignTitle
	result.CampaignResp.BannerURL = donation.CampaignBannerURL
	result.CampaignResp.TotalDonationAmount = donation.CampaignTotalAmount
	result.CampaignResp.Lembaga.ID = lembaga.ID
	result.CampaignResp.Lembaga.Name = lembaga.Name

	return
}
