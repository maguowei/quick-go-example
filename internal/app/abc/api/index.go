package api

import (
	"github.com/gin-gonic/gin"
)


func Health(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "ok",
	})
}

func Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "hello world!",
	})
}