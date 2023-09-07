package xendit

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/xendit/xendit-go/client"
)

const (
	XenditClientName = "[Xendit-Client]"
)

type IXenditClient interface {
	CreateFixedVAPayment(ctx context.Context, createVaReq *CreateFixedVAPaymentReq) (resp *VirtualAccountResp, err *Error)
	CreateQRCodePayment(ctx context.Context, createQRReq *CreateQRCodePaymentReq) (resp *QRCodeResp, err *Error)
}

type XenditClient struct {
	BaseURL      string
	ApiKey       string
	xenditClient *client.API
	HTTPClient   *http.Client
}

func NewXenditClient(xenditClient XenditClient) IXenditClient {
	xenditClient.xenditClient = client.New(xenditClient.ApiKey)
	xenditClient.HTTPClient = &http.Client{
		Timeout: 2 * time.Minute,
	}
	return &xenditClient
}

func (xc *XenditClient) newXenditRequest(ctx context.Context, path, httpMethod string, headers map[string]string, reqPayload []byte) (respPayload []byte, err *Error) {
	var request *http.Request
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	requestUrl := xc.BaseURL + path
	request, errNewRequest := http.NewRequestWithContext(ctx, httpMethod, requestUrl, bytes.NewReader(reqPayload))
	if errNewRequest != nil {
		return []byte{}, FromGoErr(http.StatusInternalServerError, errNewRequest)
	}

	request.SetBasicAuth(xc.ApiKey, "")
	request.Header.Set("Content-Type", "application/json")
	if headers != nil {
		for headerField, headerValue := range headers {
			request.Header.Set(headerField, headerValue)
		}
	}

	resp, errDoReq := xc.HTTPClient.Do(request)
	if err != nil {
		return []byte{}, FromGoErr(http.StatusInternalServerError, errDoReq)
	}
	defer resp.Body.Close()

	respPayload, errRead := io.ReadAll(resp.Body)
	if errRead != nil {
		return []byte{}, FromGoErr(http.StatusInternalServerError, errRead)
	}

	if resp.StatusCode < http.StatusOK || resp.StatusCode > 299 {
		return []byte{}, FromHTTPErr(resp.StatusCode, respPayload)
	}

	// TODO: add more log for error

	return
}
