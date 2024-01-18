package factories

import (
	"cointrading/app/config"
	"cointrading/app/presentation/backgrounds"
	"cointrading/app/presentation/controllers"
	"cointrading/app/presentation/gin"
	"cointrading/app/presentation/middleware"
	"cointrading/app/presentation/router"
)

var framework = config.FrameWork()

func NewRouter(
	sb *backgrounds.StreamIngestionDataBackground,
	oc *controllers.OrderController,
	cc *controllers.CandleController,
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

	r.POST("/sendorder", oc.SendOrder)

	// go sb.Exec()

	return r
}
