package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/e-ziswaf/eziswaf-api/internal/app/builder"
	"github.com/go-chi/chi"
)

type DonorHandler struct {
	HandlerOption
}

func NewDonorHandler(opt HandlerOption) *DonorHandler {
	return &DonorHandler{
		HandlerOption: opt,
	}
}

func (h DonorHandler) GetDonorByID(w http.ResponseWriter, r *http.Request) (err error) {
	donorID := chi.URLParam(r, "donor-id")
	if donorID == "" {
		errorMsg := builder.ErrorResponse{
			ID: "Donor ID tidak ditemukan.",
			EN: "Donor ID not found.",
		}

		errResp := builder.BuildResponse("error", errorMsg, 400)

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
		w.WriteHeader(http.StatusBadRequest)
		w.Write(respByte)
	}

	// convert string to uint64
	id, err := strconv.ParseUint(donorID, 10, 64)
	if err != nil {
		panic(err)
	}

	resp, err := h.Services.Donor.GetDonorByID(r.Context(), id)
	if err != nil {
		errorMsg := builder.ErrorResponse{
			ID: "Gagal mendapatkan detail donor. Silakan coba lagi.",
			EN: "Failed to get donor detail. Try again later.",
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
