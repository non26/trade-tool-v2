package service

import (
	service "tradetoolv2/internal/app/domain/service/binance"
	externalservices "tradetoolv2/internal/infrastructure/externalservices/binance/future"
)

type binanceFutureService struct {
	binanceFutureServiceName string
	binanceService           externalservices.IBinanceFutureExternalService
}

func NewBinanceFutureService(
	binanceFutureServiceName string,
	binanceService externalservices.IBinanceFutureExternalService,
) service.IBinanceFutureService {
	return &binanceFutureService{
		binanceFutureServiceName,
		binanceService,
	}
}
