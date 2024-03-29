package externalservices

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	entity "tradetoolv2/internal/app/domain/entity/okx"
	helper "tradetoolv2/internal/helper/okx_helper"
	okxhelper "tradetoolv2/internal/helper/okx_helper"
	"tradetoolv2/internal/infrastructure/externalservices/okx/request"
	"tradetoolv2/internal/infrastructure/externalservices/okx/response"
)

func (o *okxFutureExternalService) PlaceASinglePosition(
	e *entity.PlaceSingleOrderEntity,
) (*entity.PlaceOrderEntity, error) {

	body := &request.PlaceASinglePositionOKXServiceRequest{}
	body.ToPlaceSingleOrderRequest(e)
	if body.PosSide != "" {
		body.InstId = helper.AddInstIdUSDTSWAPPostfix(body.InstId)
	}

	_endPoint := o.okxFutureUrl.PlaceAPosition
	_url := fmt.Sprintf("%v%v", o.okxFutureUrl.OkxFutureBaseUrl.Okx1, _endPoint)
	_method := http.MethodPost
	_body, err := okxhelper.StructToJson(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(_method, _url, bytes.NewReader(_body))
	if err != nil {
		return nil, errors.New("OKX-PlaceSingleOrder Request Error: " + err.Error())
	}
	req = okxhelper.GenerateHeaders(
		req,
		_method,
		_endPoint,
		string(_body),
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
		return nil, errors.New("OKX-PlaceSingleOrder Response Error: " + err.Error())
	}
	defer res.Body.Close()

	decodeResBody := &response.CommonOKXServiceResponse{}
	err = json.NewDecoder(res.Body).Decode(decodeResBody)
	if err != nil {
		return nil, errors.New("OKX-PlaceSingleOrder Decode Error: " + err.Error())
	}

	err = helper.OkxConditionResponseError(res.StatusCode, decodeResBody.Code, decodeResBody.Message)
	if err != nil {
		return nil, err
	}

	ePlaceSingleOrder := []entity.PlaceOrderEntity{}
	ePlaceSingleOrder, err = helper.ResponseToStruct[entity.PlaceOrderEntity](ePlaceSingleOrder, decodeResBody.Data)
	if err != nil {
		return nil, err
	}
	return &ePlaceSingleOrder[0], nil
}
