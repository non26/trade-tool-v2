package service

import (
	entity "tradetoolv2/internal/app/domain/entity/okx"
	"tradetoolv2/internal/infrastructure/api/handler/okx/future/request"
	externalservices "tradetoolv2/internal/infrastructure/externalservices/okx"
)

type IOkxFutureService interface {
	PlaceAPosition(req *request.PlaceASinglePositionHandlerRequest) error
	PlaceMultiplePosition(req *request.PlaceMultiplePositionHandlerRequest) error
	SetNewLeverage(data *entity.SetLeverageFuture) (*entity.SetLeverageFuture, error)
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
