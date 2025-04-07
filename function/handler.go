package main

import (
	"github.com/gin-gonic/gin"
)

type CreateURLRequest struct {
	URL string `json:"url"`
}

type Response struct {
	ShortCode string `json:"short_code"`
}

const pathParameterName = "shortcode"
const locationHeader = "Location"

func Ding(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "dong",
	})
}
func Hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "helo world",
	})
}
