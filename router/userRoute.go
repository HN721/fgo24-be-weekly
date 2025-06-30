package router

import (
	"weekly1/controller"
	"weekly1/middleware"

	"github.com/gin-gonic/gin"
)

func userRoute(r *gin.RouterGroup) {

	r.GET("", middleware.AuthMiddleware(), controller.GetUser)
	r.GET("/:id", middleware.AuthMiddleware(), controller.GetUserById)
	r.GET("/email", middleware.AuthMiddleware(), controller.GetUserByEmail)
	r.PATCH("/upload-image", middleware.AuthMiddleware(), controller.UpdateUserImageController)

	r.PATCH("/profile", middleware.AuthMiddleware(), controller.UpdateProfileController)
	r.PATCH("/change-password", middleware.AuthMiddleware(), controller.ChangePasswordController)
}
