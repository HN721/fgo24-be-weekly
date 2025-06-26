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
	err = model.Balance(input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Success: false,
			Message: "Cannot Inset",
		})
	}
	ctx.JSON(http.StatusOK, model.Response{
		Success: true,
		Message: "Sucessfully TopUp Data ",
		Results: input,
	})
}
