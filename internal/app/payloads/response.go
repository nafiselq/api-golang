package payloads

type GetCampaignDetailByIDResp struct {
	ID             uint64  `json:"id"`
	Title          string  `json:"title"`
	BannerURL      string  `json:"banner_url"`
	DonationAmount float64 `json:"donation_amount"`
	TotalDonor     uint64  `json:"total_donor"`
	Description    string  `json:"description"`
}
