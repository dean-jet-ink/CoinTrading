package gin

import (
	"cointrading/app/presentation/router"
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ginContext struct {
	ctx *gin.Context
}

func newContext(ctx *gin.Context) router.Context {
	return &ginContext{
		ctx: ctx,
	}
}

func (g *ginContext) JSON(code int, obj any) {
	g.ctx.JSON(code, obj)
}

func (g *ginContext) String(code int, format string) {
	g.ctx.String(code, format)
}

func (g *ginContext) NoContent(code int) {
	g.ctx.Status(code)
}

func (g *ginContext) BindQuery(obj any) error {
	return g.ctx.ShouldBindQuery(obj)
}

func (g *ginContext) BindJSON(obj any) error {
	return g.ctx.ShouldBindJSON(obj)
}

func (g *ginContext) Param(key string) string {
	return g.ctx.Param(key)
}

func (g *ginContext) FormFile(name string) (*multipart.FileHeader, error) {
	return g.ctx.FormFile(name)
}

func (g *ginContext) Next() {
	g.ctx.Next()
}

func (g *ginContext) Writer() http.ResponseWriter {
	return g.ctx.Writer
}

func (g *ginContext) Request() *http.Request {
	return g.ctx.Request
}
