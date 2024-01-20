package externalservices

import (
	"net/http"
	"tradetoolv2/config"
	"tradetoolv2/internal/infrastructure/externalservices/okx/future/request"
)

type IOKXFutureExternalService interface {
	// GetLeverage() (*http.Response, error)
	SetLeverage(
		body *request.SetNewLeverageOKXServiceRequest,
	) (*http.Response, error)
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
