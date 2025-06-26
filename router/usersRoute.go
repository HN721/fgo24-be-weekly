package router

import (
	"weekly1/middleware"

	"github.com/gin-gonic/gin"
)

func userRoute(r *gin.RouterGroup) {
	r.GET("", middleware.AuthMiddleware())
}
