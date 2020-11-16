package controllers

import (
	"api.namegame.com/api/dtos"
	"api.namegame.com/services"
	"github.com/gin-gonic/gin"
)

type RoomController struct {
	RoomService services.RoomService
}

func (r *RoomController) Create(ctx *gin.Context) {
	payload := dtos.RoomCreate{}
	ctx.ShouldBind(&payload)
	fcmToken := ctx.Request.Header.Get("FCM_USER_TOKEN")
	roomState, scoreboard := r.RoomService.Create(fcmToken, payload.Username)
	gameData := dtos.GameData{Room: roomState, Scoreboard: scoreboard}
	ctx.JSON(201, gameData)
}

func (r *RoomController) Join(ctx *gin.Context) {
	roomCode := ctx.Param("room_code")
	fcmToken := ctx.Request.Header.Get("FCM_USER_TOKEN")
	payload := dtos.RoomJoin{}
	ctx.ShouldBind(&payload)
	roomState, scoreboard := r.RoomService.Join(fcmToken, roomCode, payload.Username)

	gameData := dtos.GameData{Room: roomState, Scoreboard: scoreboard}
	ctx.JSON(201, gameData)
}

func (r *RoomController) CreateHunch(ctx *gin.Context) {
	roomCode := ctx.Param("room_code")
	fcmToken := ctx.Request.Header.Get("FCM_USER_TOKEN")
	payload := dtos.HunchCreate{}
	ctx.ShouldBind(&payload)
	r.RoomService.HunchCreate(fcmToken, roomCode, payload.Hunch)
	ctx.Status(201)
}

func (r *RoomController) UpdatePlayerState(ctx *gin.Context) {
	roomCode := ctx.Param("room_code")
	fcmToken := ctx.Request.Header.Get("FCM_USER_TOKEN")
	payload := dtos.PlayerState{}
	ctx.ShouldBind(&payload)
	r.RoomService.UpdatePlayerState(fcmToken, payload.State, roomCode)
	ctx.Status(201)
}

func (r *RoomController) GetByRoomCode(ctx *gin.Context) {
	roomCode := ctx.Param("room_code")
	roomState, scoreboard := r.RoomService.GetByRoomCode(roomCode)
	gameData := dtos.GameData{Room: roomState, Scoreboard: scoreboard}
	ctx.JSON(201, gameData)
}
