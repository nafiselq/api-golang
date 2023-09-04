package payloads

type GetCampaignDetailByIDResp struct {
	ID                  uint64  `json:"id"`
	Title               string  `json:"title"`
	BannerURL           string  `json:"banner_url"`
	TotalDonationAmount float64 `json:"total_donation_amount"`
	TotalDonor          uint64  `json:"total_donor"`
	Description         string  `json:"description"`
}

type GetDonationDetailByIDResp struct {
	ID            uint64             `json:"id"`
	Amount        float64            `json:"amount"`
	Status        DonationStatusResp `json:"status"`
	PaymentMethod PaymentMethodResp  `json:"payment_method"`
	CampaignResp  CampaignResp       `json:"campaign"`
}

type DonationStatusResp struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

type CampaignResp struct {
	ID                  uint64      `json:"campaign_id"`
	Title               string      `json:"title"`
	BannerURL           string      `json:"banner_url"`
	TotalDonationAmount float64     `json:"total_donation_amount"`
	Lembaga             LembagaResp `json:"lembaga"`
}

type LembagaResp struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

type PaymentMethodResp struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}
