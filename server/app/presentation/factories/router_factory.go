package factories

import (
	"cointrading/app/config"
	"cointrading/app/presentation/backgrounds"
	"cointrading/app/presentation/controllers"
	"cointrading/app/presentation/gin"
	"cointrading/app/presentation/initialize"
	"cointrading/app/presentation/middleware"
	"cointrading/app/presentation/router"
)

var framework = config.FrameWork()

func NewRouter(
	sb *backgrounds.StreamIngestionDataBackground,
	oc *controllers.OrderController,
	cc *controllers.CandleController,
	tc *controllers.TradingConfigController,
	ini *initialize.InitExecutor,
) router.Router {
	var r router.Router

	switch framework {
	case 1:
		r = gin.NewRouter()
	default:
		r = gin.NewRouter()
	}

	r.CORS()

	r.Static("/static", "./static")

	r.Use(middleware.Recover)

	r.GET("/candles", cc.GetDataframeCandleStream)
	r.GET("trading-config", tc.GetTradingConfig)
	r.GET("/exchanges", tc.GetExchanges)
	r.GET("/symbols", tc.GetSymbols)
	r.GET("/durations", tc.GetDurations)

	r.POST("/sendorder", oc.SendOrder)

	r.PUT("/trading-config", tc.UpdateTrdingConfig)

	ini.InitializeTradingConfig()

	// go sb.Exec()

	return r
}
