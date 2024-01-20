package handler

import (
	"net/http"
	service "tradetoolv2/internal/app/service/okx/future"
	"tradetoolv2/internal/infrastructure/api/handler/okx/future/request"

	"github.com/labstack/echo/v4"
)

type IPlaceSinglePositionHandler interface {
	GetBody(c echo.Context) (*request.PlaceASinglePositionHandlerRequest, error)
	Handler(c echo.Context) error
}

type placeSinglePositionHandler struct {
	service service.IOkxFutureService
}

func NewPlaceSinglePositionHandler(
	service service.IOkxFutureService,
) IPlaceSinglePositionHandler {
	return &placeSinglePositionHandler{
		service,
	}
}

func (p *placeSinglePositionHandler) GetBody(c echo.Context) (*request.PlaceASinglePositionHandlerRequest, error) {
	body := new(request.PlaceASinglePositionHandlerRequest)
	err := c.Bind(body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (p *placeSinglePositionHandler) Handler(c echo.Context) error {
	body, err := p.GetBody(c)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			nil,
		)
	}

	err = p.service.PlaceAPosition(body)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			nil,
		)
	}
	return c.JSON(
		http.StatusOK,
		nil,
	)
}
