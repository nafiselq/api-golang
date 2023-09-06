package xendit

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

const (
	createFixedVAPayment = "/callback_virtual_accounts"
	createQRCodePayment  = "/qr_codes"
	IDRCurrency          = "IDR"
)

type CreateFixedVAPaymentReq struct {
	ExternalID     string `json:"external_id"`
	BankCode       string `json:"bank_code"`
	DonorName      string `json:"name"`
	IsSingleUse    bool   `json:"is_single_use"`
	IsClosed       bool   `json:"is_closed"`
	ExpectedAmount uint64 `json:"expected_amount"`
	ExpiresAt      string `json:"expiration_date"`
}

type VirtualAccountResp struct {
	ID             string `json:"id"`
	ExternalID     string `json:"external_id"`
	OwnerID        string `json:"owner_id"`
	BankCode       string `json:"bank_code"`
	MerchantCode   string `json:"merchant_code"`
	AccountNumber  string `json:"account_number"`
	Status         string `json:"status"`
	DonorName      string `json:"name"`
	Country        string `json:"country"`
	Currency       string `json:"currency"`
	IsSingleUse    bool   `json:"is_single_use"`
	IsClosed       bool   `json:"is_closed"`
	ExpectedAmount uint64 `json:"expected_amount"`
	ExpiresAt      string `json:"expiration_date"`
	RequestID      string `json:"-"`
}

type CreateQRCodePaymentReq struct {
	ReferenceID string `json:"reference_id"` // same as ExternalID
	Type        string `json:"type"`
	Currency    string `json:"currency"`
	Amount      uint64 `json:"amount"`
	ExpiresAt   string `json:"expires_at"`
}

type QRCodeResp struct {
	ID          string     `json:"id"`
	ReferenceID string     `json:"reference_id"`
	BusinessID  string     `json:"business_id"` // same as BusinessID
	Type        string     `json:"type"`
	Currency    string     `json:"currency"`
	Amount      uint64     `json:"amount"`
	QRString    string     `json:"qr_string"`
	Status      string     `json:"status"`
	ChannelCode string     `json:"channel_code"`
	ExpiresAt   *time.Time `json:"expires_at"`
	Created     *time.Time `json:"created"`
	Updated     *time.Time `json:"updated"`
	RequestID   string     `json:"-"`
}

func (xc *XenditClient) CreateFixedVAPayment(ctx context.Context, createVaReq *CreateFixedVAPaymentReq) (resp *VirtualAccountResp, err *Error) {
	reqByte, errMarshal := json.Marshal(createVaReq)
	if err != nil {
		return nil, FromGoErr(http.StatusInternalServerError, errMarshal)
	}

	result, requestID, err := xc.newXenditRequest(ctx, createFixedVAPayment, http.MethodPost, nil, reqByte)
	if err != nil {
		return
	}

	errUnmarshal := json.Unmarshal(result, &resp)
	if err != nil {
		return nil, FromGoErr(http.StatusInternalServerError, errUnmarshal)
	}

	resp.RequestID = requestID

	// TODO: add info log (respbody + statusCode)

	return
}

func (xc *XenditClient) CreateQRCodePayment(ctx context.Context, createQRReq *CreateQRCodePaymentReq) (resp *QRCodeResp, err *Error) {
	reqByte, errMarshal := json.Marshal(createQRReq)
	if err != nil {
		return nil, FromGoErr(http.StatusInternalServerError, errMarshal)
	}

	reqHeaders := map[string]string{
		"api-version":     "2022-07-31",
		"idempotency-key": createQRReq.ReferenceID,
	}

	result, requestID, err := xc.newXenditRequest(ctx, createQRCodePayment, http.MethodPost, reqHeaders, reqByte)
	if err != nil {
		return
	}

	errUnmarshal := json.Unmarshal(result, &resp)
	if err != nil {
		return nil, FromGoErr(http.StatusInternalServerError, errUnmarshal)
	}

	resp.RequestID = requestID

	// TODO: add info log (respbody + statusCode)

	return
}
