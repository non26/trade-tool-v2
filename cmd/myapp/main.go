package main

import (
	"fmt"
	"tradetoolv2/config"
	service "tradetoolv2/internal/app/service/binance/future"
	handler "tradetoolv2/internal/infrastructure/api/handler/binance/future"
	externalservices "tradetoolv2/internal/infrastructure/externalservices/binance/future"

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
	// binanceFutureGroup.POST("")

	app.Start(fmt.Sprintf(":%v", config.Port))
}
