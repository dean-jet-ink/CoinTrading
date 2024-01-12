package gin

import (
	"cointrading/app/config"
	"cointrading/app/presentation/router"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type ginRouter struct {
	router *gin.Engine
}

func NewRouter() router.Router {
	router := gin.New()

	return &ginRouter{
		router: router,
	}
}

func (g *ginRouter) Run(addr ...string) {
	g.router.Run(addr...)
}

func (g *ginRouter) Static(relativePath string, rootPath string) {
	g.router.Static(relativePath, rootPath)
}

func (g *ginRouter) CORS() {
	g.router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			config.FrontEndURL(),
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))
}

func (g *ginRouter) Use(middleware ...router.HandlerFunc) {
	ginHandlers := g.convertGinHandler(middleware)

	g.router.Use(ginHandlers...)
}

func (g *ginRouter) GET(relativePath string, handlers ...router.HandlerFunc) {
	ginHandlers := g.convertGinHandler(handlers)

	g.router.GET(relativePath, ginHandlers...)
}

func (g *ginRouter) POST(relativePath string, handlers ...router.HandlerFunc) {
	ginHandlers := g.convertGinHandler(handlers)

	g.router.POST(relativePath, ginHandlers...)
}

func (g *ginRouter) PUT(relativePath string, handlers ...router.HandlerFunc) {
	ginHandlers := g.convertGinHandler(handlers)

	g.router.PUT(relativePath, ginHandlers...)
}

func (g *ginRouter) DELETE(relativePath string, handlers ...router.HandlerFunc) {
	ginHandlers := g.convertGinHandler(handlers)

	g.router.DELETE(relativePath, ginHandlers...)
}

func (g *ginRouter) convertGinHandler(handlers []router.HandlerFunc) []gin.HandlerFunc {
	ginHandlers := make([]gin.HandlerFunc, len(handlers))

	for i, h := range handlers {
		ginHandler := func(c *gin.Context) {
			ginContext := newContext(c)

			h(ginContext)
		}

		ginHandlers[i] = ginHandler
	}

	return ginHandlers
}
