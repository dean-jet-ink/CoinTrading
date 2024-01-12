package router

type HandlerFunc func(Context)

type Router interface {
	Run(addr ...string)
	Static(relativePath string, rootPath string)
	CORS()
	Use(middleware ...HandlerFunc)
	GET(relativePath string, handlers ...HandlerFunc)
	POST(relativePath string, handlers ...HandlerFunc)
	PUT(relativePath string, handlers ...HandlerFunc)
	DELETE(relativePath string, handlers ...HandlerFunc)
}
