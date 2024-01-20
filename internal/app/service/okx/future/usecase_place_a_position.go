package service

import (
	"encoding/json"
	"errors"
	"fmt"
	handlerrequest "tradetoolv2/internal/infrastructure/api/handler/okx/future/request"
	"tradetoolv2/internal/infrastructure/api/handler/okx/future/response"
	"tradetoolv2/internal/infrastructure/externalservices/okx/future/request"
)

func (o *okxFutureService) PlaceAPosition(req *handlerrequest.PlaceASinglePositionHandlerRequest) error {
	body := request.PlaceASinglePositionOKXServiceRequest{
		InstId:  req.InstId,
		TdMode:  req.TdMode,
		Side:    req.Side,
		PosSide: req.PosSide,
		OrdType: req.OrdType,
		Sz:      req.Sz,
	}
	httpRes, err := o.okxExtService.PlaceASinglePosition(&body)
	if err != nil {
		return err
	}
	defer httpRes.Body.Close()

	resBody := *response.NewOkxCommonHandlerResponse(&response.PlacePositionHandlerResponse{})
	err = json.NewDecoder(httpRes.Body).Decode(resBody)
	if err != nil {
		return err
	}

	if resBody.IsCodeError() {
		return errors.New(fmt.Sprintf("OKX - Return Code %v", resBody.Code))
	}

	return nil
}
