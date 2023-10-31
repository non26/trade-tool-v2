package handler

import (
	"net/http"
	service "tradetoolv2/internal/app/domain/service/binance"
	helper "tradetoolv2/internal/helper/common_response"
	"tradetoolv2/internal/infrastructure/api/handler/binance/future/request"

	"github.com/labstack/echo/v4"
)

type ISetNewLeveragehandler interface {
	GetRequestBody(c echo.Context) (*request.SetLeverageHandlerRequest, error)
	Handler(c echo.Context) error
}

type setNewLeveragehandler struct {
	service service.IBinanceFutureService
}

func NewsetNewLeveragehandler(
	service service.IBinanceFutureService,
) ISetNewLeveragehandler {
	return &setNewLeveragehandler{
		service,
	}
}

func (h *setNewLeveragehandler) GetRequestBody(c echo.Context) (*request.SetLeverageHandlerRequest, error) {
	req := new(request.SetLeverageHandlerRequest)
	if err := c.Bind(req); err != nil {
		return nil, err
	}
	return req, nil
}

func (h *setNewLeveragehandler) Handler(c echo.Context) error {

	request, err := h.GetRequestBody(c)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			helper.CommonResponse{
				Code:    helper.FailCode,
				Message: err.Error(),
				Data:    nil,
			},
		)
	}

	err = h.service.SetNewLeverage(
		c.Request().Context(),
		request,
	)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			helper.CommonResponse{
				Code:    helper.FailCode,
				Message: err.Error(),
				Data:    nil,
			},
		)
	}

	return c.JSON(
		http.StatusOK,
		helper.CommonResponse{
			Code:    helper.SuccessCode,
			Message: "success",
			Data:    nil,
		},
	)
}
