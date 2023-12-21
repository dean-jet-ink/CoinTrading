package routers

import (
	"cointrading/app/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewGinRouter() *gin.Engine {
	router := gin.New()

	// cors制約の設定
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			config.FrontEndURL(),
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	router.Static("/static", "./static")

	return router
}
