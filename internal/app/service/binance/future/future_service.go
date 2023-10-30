package service

import (
	"context"
	"tradetoolv2/internal/infrastructure/api/handler/binance/future/request"
	externalservices "tradetoolv2/internal/infrastructure/externalservices/binance/future"
)

type IBinanceFutureService interface {
	SetNewLeverage(
		ctx context.Context,
		request *request.SetLeverageHandlerRequest) error
}

type binanceFutureService struct {
	binanceFutureServiceName string
	binanceService           externalservices.IBinanceFutureExternalService
}

func NewBinanceFutureService(
	binanceFutureServiceName string,
	binanceService externalservices.IBinanceFutureExternalService,
) IBinanceFutureService {
	return &binanceFutureService{
		binanceFutureServiceName,
		binanceService,
	}
}
