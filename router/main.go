package router

import "github.com/gin-gonic/gin"

func CombineRouter(r *gin.Engine) {
	authRoute(r.Group("/auth"))
}
