package routes

import "github.com/gin-gonic/gin"

type BaseRoute interface {
	Setup(router *gin.Engine) *gin.RouterGroup
}
