package router

import (
	"weekly1/controller"
	"weekly1/middleware"

	"github.com/gin-gonic/gin"
)

func balanceRouter(r *gin.RouterGroup) {
	r.POST("/topup", middleware.AuthMiddleware(), controller.TopUpController)
}
