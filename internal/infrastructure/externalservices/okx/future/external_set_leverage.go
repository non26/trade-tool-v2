package externalservices

import (
	"net/http"
	"tradetoolv2/internal/infrastructure/externalservices/okx/future/request"
)

func (o *okxFutureExternalService) SetLeverage(
	body *request.SetNewLeverageOKXServiceRequest,
) (*http.Response, error) {
	return nil, nil
}
