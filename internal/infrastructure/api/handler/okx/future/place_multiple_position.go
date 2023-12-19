package handler

import (
	"tradetoolv2/internal/infrastructure/api/handler/okx/future/request"

	"github.com/labstack/echo/v4"
)

type IPlaceMultiplePositionHandler interface {
	GetBody(c echo.Context) *request.PlaceMultiplePositionHandlerRequest
	Handler(c echo.Context) error
}

type placeMultiplePositionHandler struct {
}

func NewPlaceMultiplePositionHandler() IPlaceMultiplePositionHandler {
	return &placeMultiplePositionHandler{}
}

func (p *placeMultiplePositionHandler) GetBody(c echo.Context) *request.PlaceMultiplePositionHandlerRequest {
	return nil
}

func (p *placeMultiplePositionHandler) Handler(c echo.Context) error {
	return nil
}
