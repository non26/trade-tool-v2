package handler

import (
	"net/http"
	service "tradetoolv2/internal/app/service/okx/future"
	helper "tradetoolv2/internal/helper/service_response"
	"tradetoolv2/internal/infrastructure/api/handler/okx/future/request"

	"github.com/labstack/echo/v4"
)

type IPlaceSinglePositionHandler interface {
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

func (p *placeSinglePositionHandler) Handler(c echo.Context) error {
	body := new(request.PlaceASinglePositionHandlerRequest)
	err := c.Bind(body)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			&helper.CommonResponse{
				Code:    helper.FailCode,
				Message: err.Error(),
			},
		)
	}

	data, err := p.service.PlaceAPosition(body.ToPlaceSingleOrderEntity())
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			&helper.CommonResponse{
				Code:    helper.FailCode,
				Message: err.Error(),
			},
		)
	}

	return c.JSON(
		http.StatusOK,
		&helper.CommonResponse{
			Code:    helper.SuccessCode,
			Message: "Success",
			Data:    data,
		},
	)
}
