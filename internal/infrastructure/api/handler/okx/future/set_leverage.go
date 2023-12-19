package handler

import (
	"tradetoolv2/internal/infrastructure/api/handler/okx/future/request"

	"github.com/labstack/echo/v4"
)

type ISetLeverage interface {
	GetBody(c echo.Context) *request.SetNewLeverageHandlerRequest
	Handler(c echo.Context) error
}

type setLeverage struct {
}

func NewSetLeverage() ISetLeverage {
	return &setLeverage{}
}

func (s *setLeverage) GetBody(c echo.Context) *request.SetNewLeverageHandlerRequest {
	return nil
}

func (s *setLeverage) Handler(c echo.Context) error {
	return nil
}
