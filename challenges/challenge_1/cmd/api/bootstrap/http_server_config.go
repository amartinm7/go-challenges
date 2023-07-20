package bootstrap

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type httpServer struct {
	engine   *gin.Engine
	httpAddr string
}

func NewHttpServer(port int, mainContext MainContext) *httpServer {
	server := httpServer{
		engine:   gin.Default(),
		httpAddr: fmt.Sprintf(":%d", port),
	}
	server.registerRoutes(mainContext)
	return &server
}

func (server *httpServer) registerRoutes(mainContext MainContext) {
	server.engine.POST("/v1/ad", mainContext.GetHandlers()["createAdHandler"].NewGinHandler())
	server.engine.GET("/v1/ad/:id", mainContext.GetHandlers()["fetchAdHandler"].NewGinHandler())
	server.engine.GET("/v1/ad", mainContext.GetHandlers()["fetchAllAdsHandler"].NewGinHandler())
	server.engine.GET("/health", mainContext.GetHandlers()["healthCheck"].NewGinHandler())
	server.engine.GET("/", mainContext.GetHandlers()["healthCheck"].NewGinHandler())
}

func (server *httpServer) Run() error {
	fmt.Println("My Gin Server is running on httpAddr: ", server.httpAddr)
	return server.engine.Run(server.httpAddr)
}
