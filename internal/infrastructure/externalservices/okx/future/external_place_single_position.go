package externalservices

import (
	"net/http"
	"tradetoolv2/internal/infrastructure/externalservices/okx/future/request"
)

func (o *okxFutureExternalService) PlaceASinglePosition(
	body *request.PlaceASinglePositionOKXServiceRequest,
) (*http.Response, error) {
	return nil, nil
}
