package router

import (
	"weekly1/controller"
	"weekly1/middleware"

	"github.com/gin-gonic/gin"
)

func authRoute(r *gin.RouterGroup) {
	r.POST("/register", controller.RegisterController)
	r.POST("/login", controller.LoginController)
	r.GET("/get-user", middleware.AuthMiddleware(), controller.GetUser)
	r.PUT("/update-profile", middleware.AuthMiddleware(), controller.UpdateProfileController)
	r.PUT("/change-password", middleware.AuthMiddleware(), controller.ChangePasswordController)

}
