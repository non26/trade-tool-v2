package service

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	helper "tradetoolv2/internal/helper/common_response"
	"tradetoolv2/internal/infrastructure/api/handler/binance/future/request"
	"tradetoolv2/internal/infrastructure/api/handler/binance/future/response"
	BNFResponse "tradetoolv2/internal/infrastructure/externalservices/binance/future/response"
)

func (bfs *binanceFutureService) PlaceSingleOrder(
	ctx context.Context,
	request *request.PlaceSignleOrderHandlerRequest,
) (*response.PlaceSignleOrderHandlerResponse, error) {

	request.Symbol = strings.ToUpper(request.Symbol)
	request.Side = strings.ToUpper(request.Side)
	request.PositionSide = strings.ToUpper(request.PositionSide)

	bnResponse, err := bfs.binanceService.PlaceSingleOrder(
		ctx,
		request,
	)
	if err != nil {
		return nil, err
	}
	defer bnResponse.Body.Close()

	if bnResponse.StatusCode != http.StatusOK {
		bnResponseError := new(BNFResponse.ResponseBinanceFutureError)
		json.NewDecoder(bnResponse.Body).Decode(bnResponseError)
		msg := helper.FormatMessageOtherThanHttpStatus200(
			bfs.binanceFutureServiceName,
			bnResponse.StatusCode,
			bnResponseError.Code,
			bnResponseError.Message,
		)
		return nil, errors.New(msg)
	}

	res := &response.PlaceSignleOrderHandlerResponse{
		Symbol:   request.Symbol,
		Quantity: request.EntryQuantity,
	}
	return res, nil
}
