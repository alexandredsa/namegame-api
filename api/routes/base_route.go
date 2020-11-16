package routes

import "github.com/gin-gonic/gin"

//BaseRoute defines basic methods for a **Routes
type BaseRoute interface {
	Setup(router *gin.Engine) *gin.RouterGroup
}
