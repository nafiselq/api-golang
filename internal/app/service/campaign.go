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

	// TODO: Get donors count by campaign id

	resp.ID = campaign.ID
	resp.Title = campaign.Title
	resp.BannerURL = campaign.BannerURL.String
	resp.Description = campaign.Description.String
	resp.TotalDonationAmount = campaign.TotalDonationAmount
	// resp.Total = count result

	return
}
