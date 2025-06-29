package router

import (
	"weekly1/controller"
	"weekly1/middleware"

	"github.com/gin-gonic/gin"
)

func historyRoute(r *gin.RouterGroup) {
	r.POST("/createHistory", middleware.AuthMiddleware(), controller.CreateHistory)
	r.POST("/get-history", middleware.AuthMiddleware(), controller.Gethistory)
}
