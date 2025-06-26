package main

import (
	"weekly1/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.CombineRouter(r)
	r.Run()
}
