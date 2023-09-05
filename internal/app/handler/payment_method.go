package handler

import (
	"encoding/json"
	"net/http"

	"github.com/e-ziswaf/eziswaf-api/internal/app/builder"
)

type PaymentMethodHandler struct {
	HandlerOption
}

func NewPaymentMethodHandler(opt HandlerOption) *PaymentMethodHandler {
	return &PaymentMethodHandler{
		HandlerOption: opt,
	}
}

func (h PaymentMethodHandler) GetPaymentMethodList(w http.ResponseWriter, r *http.Request) (err error) {
	resp, err := h.Services.PaymentMethod.GetEnabledPaymentMethodList(r.Context())
	if err != nil {
		errorMsg := builder.ErrorResponse{
			ID: "Gagal mendapatkan list Payment Method. Silakan coba lagi.",
			EN: "Failed to get Payment Method list. Try again later.",
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
