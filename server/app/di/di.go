//+ wireinject

package di

import (
	"cointrading/app/application/interactors"
	"cointrading/app/domain/shared"
	"cointrading/app/infrastructure/factories/dbfactory"
	"cointrading/app/infrastructure/factories/tradingfactory"
	"cointrading/app/presentation/backgrounds"
	"cointrading/app/presentation/controllers"
	"cointrading/app/presentation/factories"
	"cointrading/app/presentation/initialize"
	"cointrading/app/presentation/router"
)

func Initialize() router.Router {
	// API Client
	exchange := shared.Exchange()
	tradingFactory := tradingfactory.NewTradingFactory(exchange)
	tradingAPIClient := tradingFactory.NewTradingAPIClient()

	// DB
	dbFactory := dbfactory.NewDBFactory()
	candleRepository := dbFactory.NewCandleRepository()
	tradingConfigRepository := dbFactory.NewTradingConfigRepository()

	// Background
	getRealTimeTickerUsecase := interactors.NewGetRealTimeTickerUsecase(candleRepository)
	saveCandleUsecase := interactors.NewSaveCandleUsecase(candleRepository)
	streamIngestionDataBackground := backgrounds.NewStreamIngestionDataBackground(getRealTimeTickerUsecase, saveCandleUsecase)

	sendOrderUsecase := interactors.NewOrderUsecase(tradingAPIClient)
	orderController := controllers.NewOrderController(sendOrderUsecase)

	getDataframeCandleUsecase := interactors.NewGetDataframeCandleUsecase(candleRepository)
	candleController := controllers.NewCandleController(getDataframeCandleUsecase)

	getTradingConfigUsecase := interactors.NewGetTradingConfigUsecase()
	getExchangesUsecase := interactors.NewGetExchangesUsecase()
	getSymbolsUsecase := interactors.NewGetSymbolsUsecase()
	getDurationsUsecase := interactors.NewGetDurationsUsecase()
	updateTradingConfigUsecase := interactors.NewUpdateTradingConfigUsecase(tradingConfigRepository)
	tradingConfigController := controllers.NewTradingConfigController(getTradingConfigUsecase, getExchangesUsecase, getSymbolsUsecase, getDurationsUsecase, updateTradingConfigUsecase)

	initializeTradingConfigUsecase := interactors.NewInitializeTradingConfigUsecase(tradingConfigRepository)
	initExcutor := initialize.NewInitExecutor(initializeTradingConfigUsecase)

	routerFactory := factories.NewRouter(
		streamIngestionDataBackground,
		orderController,
		candleController,
		tradingConfigController,
		initExcutor,
	)

	return routerFactory
}
