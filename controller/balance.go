package controller

import (
	"net/http"
	"weekly1/model"

	"github.com/gin-gonic/gin"
)

func BalanceController(ctx *gin.Context) {
	var input model.Balances
	err := ctx.ShouldBindJSON(&input)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Success: false,
			Message: "Cannot Insert Data or Invalid Data",
		})
		return
	}
	userId, _ := ctx.Get("userID")

	err = model.Balance(userId.(int), input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Success: false,
			Message: "Cannot Inset",
		})
	}
	ctx.JSON(http.StatusOK, model.Response{
		Success: true,
		Message: "Sucessfully Create Balance  Data ",
		Results: input,
	})
}
func TopUpController(ctx *gin.Context) {
	var top model.TopUps
	userId, _ := ctx.Get("userID")

	err := ctx.ShouldBindJSON(&top)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Success: false,
			Message: "Cannot Insert Data or Invalid Data",
		})
		return
	}
	err = model.TopUp(userId.(int), top)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Success: false,
			Message: "Cannot Inset",
		})
	}
	ctx.JSON(http.StatusOK, model.Response{
		Success: true,
		Message: "Sucessfully TopUp Data ",
		Results: top,
	})
}
