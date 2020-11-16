package controllers

import (
	"api.namegame.com/api/dtos"
	"github.com/gin-gonic/gin"
)

type RoomController struct {
}

//New func to calculate new Withdrawal
func (r *RoomController) Create(ctx *gin.Context) {
	payload := dtos.RoomCreate{}
	ctx.ShouldBind(&payload)

	banknotes, err := r.BanknoteDataService.Withdrawal(payload.Value)

	if err != nil {
		ctx.JSON(422, err)
		return
	}

	ctx.JSON(201, banknotes)
}
