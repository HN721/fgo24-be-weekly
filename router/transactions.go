package router

import (
	"weekly1/controller"
	"weekly1/middleware"

	"github.com/gin-gonic/gin"
)

func transactionRoute(r *gin.RouterGroup) {
	r.POST("", middleware.AuthMiddleware(), controller.CreateTransactionsController)
	r.GET("", middleware.AuthMiddleware(), controller.GetUserTransactionController)

}
