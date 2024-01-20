package main

import (
	"fmt"
	"net/http"
	"tradetoolv2/config"
	service "tradetoolv2/internal/app/service/binance/future"
	okxservice "tradetoolv2/internal/app/service/okx/future"
	handler "tradetoolv2/internal/infrastructure/api/handler/binance/future"
	okxhandler "tradetoolv2/internal/infrastructure/api/handler/okx/future"
	externalservices "tradetoolv2/internal/infrastructure/externalservices/binance/future"
	okxexternalservices "tradetoolv2/internal/infrastructure/externalservices/okx/future"

	"github.com/labstack/echo/v4"
)

func BinanceFutureGroup(g *echo.Group) {}

func BinanceSpotGroup() {}

func main() {
	config, err := config.ReadConfig()
	if err != nil {
		panic(err.Error())
	}

	app := echo.New()

	app.GET("/", func(c echo.Context) error {
		type HealthCheck struct {
			Message string `json:"message"`
		}
		return c.JSON(
			http.StatusOK,
			&HealthCheck{
				Message: "success",
			},
		)
	})

	// binance future
	binanceFutureGroup := app.Group("/" + config.ServiceName.BinanceFuture)
	binanceFutureExternalService := externalservices.NewBinanceFutureExternalService(
		&config.BinanceFutureUrl,
		&config.Secrets,
	)
	binanceFutureService := service.NewBinanceFutureService(
		config.ServiceName.BinanceFuture,
		binanceFutureExternalService,
	)

	setNewLeveragehandler := handler.NewsetNewLeveragehandler(
		binanceFutureService,
	)
	binanceFutureGroup.POST("/set-new-leverage", setNewLeveragehandler.Handler)

	placeSingleOrderHandler := handler.NewPlaceSinglerOrderHandler(
		binanceFutureService,
	)
	binanceFutureGroup.POST("/place-a-order", placeSingleOrderHandler.Handler)

	// okx future
	okxFutureGroup := app.Group("/" + config.ServiceName.OKXFuture)
	okxFutureExternalService := okxexternalservices.NewOKXFutureExternalService(
		&config.OkxFutureUrl,
		&config.Secrets,
		config.Env,
	)
	okxService := okxservice.NewOkxService(
		okxFutureExternalService,
	)

	okx1 := okxhandler.NewPlaceSinglePositionHandler(okxService)
	okxFutureGroup.POST(
		"/placeSinglePosition",
		okx1.Handler,
	)

	app.Start(fmt.Sprintf(":%v", config.Port))
}
