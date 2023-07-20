package handler

import "github.com/gin-gonic/gin"

type GinHandler interface {
	NewGinHandler() gin.HandlerFunc
}
