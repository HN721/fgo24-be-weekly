package controller

import (
	"net/http"
	"weekly1/model"

	"github.com/gin-gonic/gin"
)

func LoginController(ctx *gin.Context) {
	var input model.Profile

	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Success: false,
			Message: "Input tidak valid",
		})
		return
	}

	user, err := model.Login(input)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, model.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, model.Response{
		Success: true,
		Message: "Login berhasil",
		Results: user,
	})
}
func RegisterController(ctx *gin.Context) {
	var input model.Profile
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Success: false,
			Message: "Input Tidak Valid",
		})
		return
	}
	err = model.Register(input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Success: false,
			Message: "Input Tidak Valid",
		})
		return
	}
	ctx.JSON(http.StatusOK, model.Response{
		Success: true,
		Message: "Sucessfully",
		Results: input,
	})
}
