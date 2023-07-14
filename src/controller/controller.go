package controller

import (
	"urlshortner/src/models"
	"urlshortner/src/services"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	c.JSON(200, "Welcome to url shortner")
}

func ShortUrl(c *gin.Context) {
	var url models.Url
	url.LongUrl = c.Query("longurl")
	url.ShortUrl = "www." + services.ShortTheUrl(url.LongUrl) + ".com"
	c.JSON(200, gin.H{
		"long_url":  url.LongUrl,
		"short_url": "www." + url.ShortUrl + ".com",
	})

}
