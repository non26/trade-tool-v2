package externalservices

import (
	"net/http"
	"tradetoolv2/internal/infrastructure/externalservices/okx/future/request"
)

func (o *okxFutureExternalService) PlaceMultiplePosition(
	body *request.PlaceMultiplePositionOKXServiceRequest,
) (*http.Response, error) {
	return nil, nil
}
