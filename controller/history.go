package controller

import (
	"net/http"
	"weekly1/model"

	"github.com/gin-gonic/gin"
)

func CreateHistory(ctx *gin.Context) {
	var input model.Historymodel

	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Success: false,
			Message: "Invalid Input",
		})
	}
	err = model.CreateHistory(input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Success: false,
			Message: "Cannot Insert Data",
		})
	}
	ctx.JSON(http.StatusOK, model.Response{
		Success: true,
		Message: "Sucessfully Addede History Data",
		Results: input,
	})
}
func Gethistory(ctx *gin.Context) {}
