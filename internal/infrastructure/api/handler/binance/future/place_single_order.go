package handler

import (
	"net/http"
	service "tradetoolv2/internal/app/domain/service/binance"
	helper "tradetoolv2/internal/helper/common_response"
	"tradetoolv2/internal/infrastructure/api/handler/binance/future/request"

	"github.com/labstack/echo/v4"
)

type IPlaceSingleOrderHandler interface {
	GetRequestBody(c echo.Context) (*request.PlaceSignleOrderHandlerRequest, error)
	Handler(c echo.Context) error
}

type placeSinglerOrderHandler struct {
	service service.IBinanceFutureService
}

func NewPlaceSinglerOrderHandler(
	service service.IBinanceFutureService,
) IPlaceSingleOrderHandler {
	return &placeSinglerOrderHandler{
		service,
	}
}

func (h *placeSinglerOrderHandler) GetRequestBody(c echo.Context) (*request.PlaceSignleOrderHandlerRequest, error) {
	req := new(request.PlaceSignleOrderHandlerRequest)
	if err := c.Bind(req); err != nil {
		return nil, err
	}
	return req, nil
}

func (h *placeSinglerOrderHandler) Handler(c echo.Context) error {

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

	res, err := h.service.PlaceSingleOrder(
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
			Data:    res,
		},
	)
}
