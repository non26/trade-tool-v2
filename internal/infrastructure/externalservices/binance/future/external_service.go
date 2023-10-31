package externalservices

import (
	"context"
	"net/http"
	"tradetoolv2/config"
	"tradetoolv2/internal/infrastructure/api/handler/binance/future/request"
)

type IBinanceFutureExternalService interface {
	SetNewLeverage(
		ctx context.Context,
		request *request.SetLeverageHandlerRequest) (*http.Response, error)
	PlaceSingleOrder(
		ctx context.Context,
		request *request.PlaceSignleOrderHandlerRequest,
	) (*http.Response, error)
}

type binanceFutureExternalService struct {
	binanceFutureUrl *config.BinanceFutureUrl
	secrets          *config.Secrets
}

func NewBinanceFutureExternalService(
	binanceFutureUrl *config.BinanceFutureUrl,
	secrets *config.Secrets,
) IBinanceFutureExternalService {
	return &binanceFutureExternalService{
		binanceFutureUrl,
		secrets,
	}
}
