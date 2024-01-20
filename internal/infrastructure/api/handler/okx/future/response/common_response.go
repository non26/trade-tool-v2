package response

type OkxCommonHandlerResponse struct {
	Code    string      `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
	InTime  string      `json:"inTime"`
	OutTime string      `json:"outTime"`
}

func NewOkxCommonHandlerResponse(data interface{}) *OkxCommonHandlerResponse {
	o := OkxCommonHandlerResponse{}
	o.Data = data
	return &o
}

func (o *OkxCommonHandlerResponse) IsCodeError() (isErr bool) {
	switch o.Code {
	case "0":
		isErr = false
	default:
		isErr = true
	}
	return isErr
}
