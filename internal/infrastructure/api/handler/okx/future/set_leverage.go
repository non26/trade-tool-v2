package handler

import (
	"net/http"

	service "tradetoolv2/internal/app/service/okx/future"
	helper "tradetoolv2/internal/helper/service_response"
	"tradetoolv2/internal/infrastructure/api/handler/okx/future/request"

	"github.com/labstack/echo/v4"
)

type ISetLeverage interface {
	Handler(c echo.Context) error
}

type setLeverage struct {
	okxservice service.IOkxFutureService
}

func NewSetLeverage(
	okxservice service.IOkxFutureService,
) ISetLeverage {
	return &setLeverage{
		okxservice,
	}
}
func (s *setLeverage) Handler(c echo.Context) error {
	body := new(request.SetNewLeverageHandlerRequest)
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

	data, err := s.okxservice.SetNewLeverage(body.ToSetLeverageEntity())
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			&helper.CommonResponse{
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
			Message: "Success",
			Data:    data,
		},
	)
}
