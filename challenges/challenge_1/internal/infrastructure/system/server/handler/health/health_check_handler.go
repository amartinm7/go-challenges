package health

import (
	"challenges/challenge_1/internal/infrastructure/kit/server/handler"
	"github.com/gin-gonic/gin"
	"net/http"
)

type healthCheckHandler struct {
}

func NewHealthCheckHandler() handler.GinHandler {
	return &healthCheckHandler{}
}
func (handler *healthCheckHandler) NewGinHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, healthCheckResponse{Status: "OK", Message: "Service running!"})
	}
}

type healthCheckResponse struct {
	Status  string
	Message string
}
