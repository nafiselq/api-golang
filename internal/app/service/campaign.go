package service

import (
	"context"

	"github.com/e-ziswaf/eziswaf-api/internal/app/payloads"
)

type ICampaignService interface {
	GetCampaignByID(ctx context.Context, campaignID uint64) (resp payloads.GetCampaignDetailByIDResp, err error)
}

type CampaignService struct {
	opt Option
}

func NewCampaignService(opt Option) ICampaignService {
	return &CampaignService{
		opt: opt,
	}
}

func (cs *CampaignService) GetCampaignByID(ctx context.Context, campaignID uint64) (resp payloads.GetCampaignDetailByIDResp, err error) {
	campaign, err := cs.opt.Repository.Campaign.GetCampaignByID(ctx, campaignID)
	if err != nil {
		// TODO: Add log & split error between not found and other error
		return
	}

	lembaga, err := cs.opt.Repository.Lembaga.GetLembagaByID(ctx, campaign.LembagaID)
	if err != nil {
		// TODO: Add log & split error between not found and other error
		return
	}

	resp.ID = campaign.ID
	resp.Title = campaign.Title
	resp.BannerURL = campaign.BannerURL.String
	resp.Description = campaign.Description.String
	resp.TotalDonationAmount = campaign.TotalDonationAmount
	resp.Lembaga.ID = lembaga.ID
	resp.Lembaga.Name = lembaga.Name

	return
}
