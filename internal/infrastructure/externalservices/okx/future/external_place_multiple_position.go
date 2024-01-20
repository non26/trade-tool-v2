package externalservices

import (
	"errors"
	"fmt"
	"net/http"
	okxhelper "tradetoolv2/internal/helper/okx_helper"
	"tradetoolv2/internal/infrastructure/externalservices/okx/future/request"
)

func (o *okxFutureExternalService) PlaceMultiplePosition(
	body *request.PlaceMultiplePositionOKXServiceRequest,
) (*http.Response, error) {
	_endPoint := o.okxFutureUrl.PlaceMultiPosition
	_url := fmt.Sprintf("%v%v", o.okxFutureUrl.OkxFutureBaseUrl.Okx1, _endPoint)
	_method := http.MethodPost
	_body, err := okxhelper.StructToJson(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(_method, _url, _body)
	if err != nil {
		return nil, errors.New("OKX-PlaceMultipleOrder Request Error: " + err.Error())
	}
	req = okxhelper.GenerateHeaders(
		req,
		_method,
		_endPoint,
		okxhelper.ToQueryParameter(body),
		o.env,
		o.secret,
	)

	transport := &http.Transport{}
	client := &http.Client{
		Transport: transport,
	}
	var res *http.Response
	res, err = client.Do(req)
	if err != nil {
		return nil, errors.New("OKX-PlaceMultipleOrder Response Error: " + err.Error())
	}

	return res, nil
}
