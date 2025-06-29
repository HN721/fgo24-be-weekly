package controller

import (
	"net/http"
	"weekly1/model"

	"github.com/gin-gonic/gin"
)

func CreateTransactionsController(ctx *gin.Context) {
	var input model.Transactions

	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Success: false,
			Message: "Cannot Insert Data or Invalid Data",
		})
		return
	}
	userId, _ := ctx.Get("userID")
	err = model.CreateTransactions(userId.(int), input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Success: false,
			Message: "Cannot Inset",
		})
	}
	ctx.JSON(http.StatusOK, model.Response{
		Success: true,
		Message: "Sucessfully Transfer  ",
		Results: input,
	})
}
func GetUserTransactionController(ctx *gin.Context) {
	userIDVal, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, model.Response{
			Success: false,
			Message: "Unauthorized: user ID not found",
		})
		return
	}

	userID, ok := userIDVal.(int)
	if !ok {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Success: false,
			Message: "Invalid user ID format",
		})
		return
	}

	transactions, err := model.GetTransactionByUserId(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.Response{
			Success: false,
			Message: "Failed to get transactions: " + err.Error(),
		})
		return
	}

	if len(transactions) == 0 {
		ctx.JSON(http.StatusNotFound, model.Response{
			Success: false,
			Message: "No transactions found",
		})
		return
	}

	ctx.JSON(http.StatusOK, model.Response{
		Success: true,
		Message: "Successfully retrieved transactions",
		Results: transactions,
	})
}
