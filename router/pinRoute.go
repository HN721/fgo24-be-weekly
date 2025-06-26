package router

import (
	"weekly1/controller"
	"weekly1/middleware"

	"github.com/gin-gonic/gin"
)

func pinRoute(r *gin.RouterGroup) {
	r.POST("", middleware.AuthMiddleware(), controller.PinController)
}
