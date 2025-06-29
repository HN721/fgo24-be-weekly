package router

import (
	"weekly1/controller"
	"weekly1/middleware"

	"github.com/gin-gonic/gin"
)

func transactionRoute(r *gin.RouterGroup) {
	r.POST("/createTransaction", middleware.AuthMiddleware(), controller.CreateTransactionsController)
	r.GET("/findTransaction", middleware.AuthMiddleware(), controller.GetUserTransactionController)

}
