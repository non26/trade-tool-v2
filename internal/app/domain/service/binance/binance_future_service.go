package service

import (
	"context"
	"tradetoolv2/internal/infrastructure/api/handler/binance/future/request"
	"tradetoolv2/internal/infrastructure/api/handler/binance/future/response"
)

type IBinanceFutureService interface {
	SetNewLeverage(
		ctx context.Context,
		request *request.SetLeverageHandlerRequest) error
	PlaceSingleOrder(
		ctx context.Context,
		request *request.PlaceSignleOrderHandlerRequest,
	) (*response.PlaceSignleOrderHandlerResponse, error)
}
