package handler

import "github.com/labstack/echo/v4"

type IGetLeverageHandler interface {
	Handler(c echo.Context) error
}

type getLeverageHandler struct {
	// service
}

func NreGetLeverageHandler() IGetLeverageHandler {
	return &getLeverageHandler{}
}

func (g *getLeverageHandler) Handler(c echo.Context) error {
	return nil
}
