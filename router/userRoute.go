package router

import (
	"weekly1/controller"
	"weekly1/middleware"

	"github.com/gin-gonic/gin"
)

func userRoute(r *gin.RouterGroup) {

	r.GET("", middleware.AuthMiddleware(), controller.GetUser)
	r.PATCH("/profile", middleware.AuthMiddleware(), controller.UpdateProfileController)
	r.PATCH("/change-password", middleware.AuthMiddleware(), controller.ChangePasswordController)
}
