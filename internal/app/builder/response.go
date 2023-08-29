package builder

import "encoding/json"

// Response API response
type Response struct {
}

type SuccessResp struct {
	StatusCode int         `json:"status_code"`
	Status     string      `json:"status"`
	Data       interface{} `json:"data,omitempty"`
}

type ErrorResp struct {
	StatusCode int           `json:"status_code"`
	Status     string        `json:"status"`
	Error      ErrorResponse `json:"error,omitempty"`
}

// ErrorResponse error response detail
type ErrorResponse struct {
	ID string `json:"id"`
	EN string `json:"en"`
}

// BuildResponse build api response
func BuildResponse(respType string, data interface{}, statusCode int) interface{} {
	if respType == "success" {
		resp := SuccessResp{}
		resp.Status = "success"
		resp.StatusCode = statusCode
		resp.Data = data

		return resp
	} else {
		resp := ErrorResp{}
		resp.Status = "error"
		resp.StatusCode = statusCode
		resp.Error = data.(ErrorResponse)

		return resp
	}
}

// ErrInternalServer error internal server
var ErrInternalServer = ErrorResp{
	StatusCode: 500,
	Status:     "error",
	Error: ErrorResponse{
		ID: "Terjadi kesalahan. Silakan coba lagi.",
		EN: "Ups, something wrong with the server. Please try again later",
	},
}

// GenerateInternalServerError generate error resp internal server error
func GenerateInternalServerError() []byte {
	byteData, _ := json.Marshal(ErrInternalServer)

	return byteData
}
