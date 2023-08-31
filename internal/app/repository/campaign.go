package repository

import (
	"context"

	"github.com/e-ziswaf/eziswaf-api/internal/app/model"
)

type ICampaignRepository interface {
	GetCampaignByID(ctx context.Context, campaignID uint64) (campaign model.Campaign, err error)
}

type CampaignRepository struct {
	opt Option
}

func NewCampaignRepository(opt Option) ICampaignRepository {
	return &CampaignRepository{
		opt: opt,
	}
}

func (cr *CampaignRepository) GetCampaignByID(ctx context.Context, campaignID uint64) (campaign model.Campaign, err error) {
	result := cr.opt.DbPostgre.WithContext(ctx).First(&campaign, campaignID)
	err = result.Error
	return
}
