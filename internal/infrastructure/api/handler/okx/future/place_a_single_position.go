package handler

import (
	"tradetoolv2/internal/infrastructure/api/handler/okx/future/request"

	"github.com/labstack/echo/v4"
)

type IPlaceSinglePositionHandler interface {
	GetBody(c echo.Context) *request.PlaceASinglePositionHandlerRequest
	Handler(c echo.Context) error
}

type placeSinglePositionHandler struct {
}

func NewPlaceSinglePositionHandler() IPlaceSinglePositionHandler {
	return &placeSinglePositionHandler{}
}

func (p *placeSinglePositionHandler) GetBody(c echo.Context) *request.PlaceASinglePositionHandlerRequest {
	return nil
}

func (p *placeSinglePositionHandler) Handler(c echo.Context) error {
	return nil
}
