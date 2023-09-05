package xendit

import (
	"context"
	"fmt"

	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/client"
	"github.com/xendit/xendit-go/qrcode"
	"github.com/xendit/xendit-go/virtualaccount"
)

type IXenditClient interface {
}

type XenditClient struct {
	ApiKey       string
	xenditClient *client.API
}

func NewXenditClient(xendit XenditClient) IXenditClient {
	xendit.xenditClient = client.New(xendit.ApiKey)
	return &xendit
}

func (xc *XenditClient) CreateFixedVAPayment(ctx context.Context, createVaReq *virtualaccount.CreateFixedVAParams) (resp *xendit.VirtualAccount, err error) {
	resp, errCreateFixedVA := xc.xenditClient.VirtualAccount.CreateFixedVAWithContext(ctx, createVaReq)
	if errCreateFixedVA != nil {
		err = fmt.Errorf("Error Create Fixed VA payment from Xendit | %d | %s", errCreateFixedVA.GetErrorCode(), errCreateFixedVA.Error())
		return
	}

	return
}

func (xc *XenditClient) CreateQRCodePayment(ctx context.Context, createQRReq *qrcode.CreateQRCodeParams) (resp *xendit.QRCode, err error) {
	resp, errCreateQRPayment := xc.xenditClient.QRCode.CreateQRCodeWithContext(ctx, createQRReq)
	if errCreateQRPayment != nil {
		err = fmt.Errorf("Error Create QR payment from Xendit | %d | %s", errCreateQRPayment.GetErrorCode(), errCreateQRPayment.Error())
		return
	}

	return
}
