package service

import (
	"tradetoolv2/internal/infrastructure/api/handler/okx/future/request"
	externalservices "tradetoolv2/internal/infrastructure/externalservices/okx/future"
)

type IOkxFutureService interface {
	PlaceAPosition(req *request.PlaceASinglePositionHandlerRequest) error
	PlaceMultiplePosition(req *request.PlaceMultiplePositionHandlerRequest) error
	SetNewLeverage(req *request.SetNewLeverageHandlerRequest) error
}

type okxFutureService struct {
	okxExtService externalservices.IOKXFutureExternalService
}

func NewOkxService(
	okxExtService externalservices.IOKXFutureExternalService,
) IOkxFutureService {
	return &okxFutureService{
		okxExtService,
	}
}
