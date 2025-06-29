package router

import "github.com/gin-gonic/gin"

func CombineRouter(r *gin.Engine) {
	authRoute(r.Group("/auth"))
	pinRoute(r.Group("/auth/pin"))
	balanceRouter(r.Group("/balance"))
	transactionRoute(r.Group("/transaction"))
	historyRoute(r.Group("/historys"))
}
