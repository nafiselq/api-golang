package repository

import (
	"context"

	"github.com/e-ziswaf/eziswaf-api/internal/app/model"
)

type IDonationRepository interface {
	GetDonationDetailByID(ctx context.Context, donationID uint64) (donationDetail model.DonationDetail, err error)
}

type DonationRepository struct {
	opt Option
}

func NewDonationRepository(opt Option) IDonationRepository {
	return &DonationRepository{
		opt: opt,
	}
}

func (dr *DonationRepository) GetDonationDetailByID(ctx context.Context, donationID uint64) (donationDetail model.DonationDetail, err error) {
	// result := dr.opt.DbPostgre.WithContext(ctx).First(&donationDetail, donationID) // join
	// if result.Error != nil {
	// 	err = result.Error
	// }

	query := `SELECT d.id, d.amount, d.status_id, d.campaign_id, d.payment_method_id, 
       	ds.name donation_status_name,  
       	pm.name payment_method_name, 
       	c.title campaign_title, 
       	c.banner_url campaign_banner_url, 
       	c.total_donation_amount campaign_total_amount, 
       	c.lembaga_id campaign_lembaga_id 
		from donation d join donation_status ds on ds.id = d.status_id
		join payment_method pm on pm.id = d.payment_method_id
		join campaign c on c.id = d.campaign_id
		where d.id = ?`

	result := dr.opt.DbPostgre.WithContext(ctx).Raw(query, donationID).Scan(&donationDetail)
	if result.Error != nil {
		err = result.Error
		return
	}

	return
}
