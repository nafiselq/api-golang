package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/e-ziswaf/eziswaf-api/internal/app/builder"
	"github.com/go-chi/chi"
)

type CampaignHandler struct {
	HandlerOption
}

func NewCampaignHandler(opt HandlerOption) *CampaignHandler {
	return &CampaignHandler{
		HandlerOption: opt,
	}
}

func (h CampaignHandler) GetCampaignByID(w http.ResponseWriter, r *http.Request) (err error) {
	ctx := r.Context()
	campaignIDStr := chi.URLParam(r, "campaign_id")
	campaignID, err := strconv.ParseUint(campaignIDStr, 10, 64)
	if err != nil {
		panic(err) // need better error handling other thna panic
	}

	resp, err := h.Services.Campaign.GetCampaignByID(ctx, campaignID)
	if err != nil {
		errorMsg := builder.ErrorResponse{
			ID: "Gagal mendapatkan detail Campaign. Silakan coba lagi.",
			EN: "Failed to get Campaign detail. Try again later.",
		}

		errResp := builder.BuildResponse("error", errorMsg, 500)

		respByte, err := json.Marshal(errResp)
		if err != nil {
			// TODO: add logger

			errMsg := builder.GenerateInternalServerError()
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(errMsg)

			return err
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(respByte)

		return err
	}

	response := builder.BuildResponse("success", resp, 200)
	respByte, err := json.Marshal(response)
	if err != nil {
		// TODO: add logger

		errMsg := builder.GenerateInternalServerError()
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(errMsg)

		return err
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respByte)

	return
}
