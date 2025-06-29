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
		return
	}
	err = model.CreateHistory(input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Success: false,
			Message: "Cannot Insert Data",
		})
		return
	}
	ctx.JSON(http.StatusOK, model.Response{
		Success: true,
		Message: "Sucessfully Addede History Data",
		Results: input,
	})
}
func Gethistory(ctx *gin.Context) {
	histories, err := model.GetAllHistory()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.Response{
			Success: false,
			Message: "Fail to retrieve history: " + err.Error(),
		})
		return
	}

	if len(histories) == 0 {
		ctx.JSON(http.StatusNotFound, model.Response{
			Success: false,
			Message: "No history found",
		})
		return
	}

	ctx.JSON(http.StatusOK, model.Response{
		Success: true,
		Message: "Successfully retrieved history",
		Results: histories,
	})
}
