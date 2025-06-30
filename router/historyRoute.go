package router

import (
	"weekly1/controller"
	"weekly1/middleware"

	"github.com/gin-gonic/gin"
)

func historyRoute(r *gin.RouterGroup) {
	r.POST("", middleware.AuthMiddleware(), controller.CreateHistory)
	r.GET("", middleware.AuthMiddleware(), controller.Gethistory)
}
