package externalservices

import (
	"net/http"
	"tradetoolv2/config"
	entity "tradetoolv2/internal/app/domain/entity/okx"
	"tradetoolv2/internal/infrastructure/externalservices/okx/request"
)

type IOKXFutureExternalService interface {
	// GetLeverage() (*http.Response, error)
	SetLeverage(
		body *entity.SetLeverageFuture,
	) (*entity.SetLeverageFuture, error)
	PlaceASinglePosition(
		body *request.PlaceASinglePositionOKXServiceRequest,
	) (*http.Response, error)
	PlaceMultiplePosition(
		body *request.PlaceMultiplePositionOKXServiceRequest,
	) (*http.Response, error)
}

type okxFutureExternalService struct {
	okxFutureUrl *config.OkxFutureUrl
	secret       *config.Secrets
	env          string
}

func NewOKXFutureExternalService(
	okxFutureUrl *config.OkxFutureUrl,
	secret *config.Secrets,
	env string,
) IOKXFutureExternalService {
	return &okxFutureExternalService{
		okxFutureUrl,
		secret,
		env,
	}
}
