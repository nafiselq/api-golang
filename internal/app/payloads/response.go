package payloads

type GetCampaignDetailByIDResp struct {
	ID                  uint64      `json:"id"`
	Title               string      `json:"title"`
	BannerURL           string      `json:"banner_url"`
	TotalDonationAmount float64     `json:"total_donation_amount"`
	Description         string      `json:"description"`
	Lembaga             LembagaResp `json:"lembaaga"`
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
	ID          uint64  `json:"id"`
	Name        string  `json:"name"`
	BankCode    string  `json:"bank_code,omitempty"`
	Type        string  `json:"type,omitempty"`
	IsEnabled   bool    `json:"is_enabled,omitempty"`
	FixedFee    float32 `json:"fixed_fee,omitempty"`
	VariableFee float32 `json:"variable_fee,omitempty"`
	Logo        string  `json:"logo,omitempty"`
}
