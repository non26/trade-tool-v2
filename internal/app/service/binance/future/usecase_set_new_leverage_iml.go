package service

import (
	"context"
	"encoding/json"
	"errors"
	helper "tradetoolv2/internal/helper/service_response"
	"tradetoolv2/internal/infrastructure/api/handler/binance/future/request"
	"tradetoolv2/internal/infrastructure/externalservices/binance/future/response"
)

func (bfs *binanceFutureService) SetNewLeverage(
	ctx context.Context,
	request *request.SetLeverageHandlerRequest) error {
	bnResponse, err := bfs.binanceService.SetNewLeverage(
		ctx,
		request,
	)
	if err != nil {
		return err
	}
	defer bnResponse.Body.Close()

	if bnResponse.StatusCode != 200 {
		bnResponseError := new(response.ResponseBinanceFutureError)
		json.NewDecoder(bnResponse.Body).Decode(bnResponseError)
		msg := helper.FormatMessageOtherThanHttpStatus200(
			bfs.binanceFutureServiceName,
			bnResponse.StatusCode,
			bnResponseError.Code,
			bnResponseError.Message,
		)
		return errors.New(msg)

	}
	return nil
}
