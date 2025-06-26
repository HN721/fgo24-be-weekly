package router

import (
	"weekly1/controller"

	"github.com/gin-gonic/gin"
)

func authRoute(r *gin.RouterGroup) {
	r.POST("/register", controller.RegisterController)
	r.POST("/login", controller.LoginController)
}
