package routes

import (
	"api.namegame.com/api/controllers"
	"github.com/gin-gonic/gin"
)

type RoomRoute struct {
	RoomController controllers.RoomController
}

func (r RoomRoute) Setup(router *gin.Engine) *gin.RouterGroup {
	routes := router.Group("/rooms")
	{
		routes.GET("/:room_code", r.RoomController.GetByRoomCode)
		routes.POST("/create", r.RoomController.Create)
		routes.POST("/join/:room_code", r.RoomController.Join)
		routes.POST("/hunches/:room_code", r.RoomController.CreateHunch)
		routes.PUT("/players/:room_code/me", r.RoomController.UpdatePlayerState)
	}

	return routes
}
