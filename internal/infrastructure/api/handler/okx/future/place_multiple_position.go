package handler

import (
	"net/http"
	service "tradetoolv2/internal/app/service/okx/future"
	helper "tradetoolv2/internal/helper/service_response"
	"tradetoolv2/internal/infrastructure/api/handler/okx/future/request"

	"github.com/labstack/echo/v4"
)

type IPlaceMultiplePositionHandler interface {
	Handler(c echo.Context) error
}

type placeMultiplePositionHandler struct {
	service service.IOkxFutureService
}

func NewPlaceMultiplePositionHandler(
	service service.IOkxFutureService,
) IPlaceMultiplePositionHandler {
	return &placeMultiplePositionHandler{
		service,
	}
}

func (p *placeMultiplePositionHandler) GetBody(c echo.Context) *request.PlaceMultiplePositionHandlerRequest {
	return nil
}

func (p *placeMultiplePositionHandler) Handler(c echo.Context) error {

	body := new(request.PlaceMultiplePositionHandlerRequest)
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

	err = p.service.PlaceMultiplePosition(body)
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
		helper.CommonResponse{
			Code:    helper.SuccessCode,
			Message: "Success",
			Data:    nil,
		},
	)
}
