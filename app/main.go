package main

import (
	"cointrading/app/application/interactors"
	"cointrading/app/config"
	"cointrading/app/presentation/backgrounds"
	"cointrading/app/presentation/routers"
	"fmt"
)

func main() {
	config.LoggingSettings(config.LogFileName())

	getRealTimeTickerUsecase := interactors.NewGetRealTimeTickerUsecase()

	background := backgrounds.NewGetRealTimeTickerBackground(getRealTimeTickerUsecase)

	background.Exec()

	router := routers.NewGinRouter()

	router.Run(fmt.Sprintf(":%v", config.Port()))
}
