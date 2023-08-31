package handler

import "net/http"

type CampaignHandler struct {
	HandlerOption
}

func NewCampaignHandler(opt HandlerOption) *CampaignHandler {
	return &CampaignHandler{
		HandlerOption: opt,
	}
}

func (h CampaignHandler) GetCampaignByID(w http.ResponseWriter, r *http.Request) (err error) {

	return
}