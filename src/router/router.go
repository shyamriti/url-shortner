package router

import (
	"urlshortner/src/controller"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.GET("/home", controller.Home)
	r.GET("/short", controller.ShortUrl)
	return r
}
