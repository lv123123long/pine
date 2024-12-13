package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lv123123long/pine/controllers"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/login", controllers.Login)
	r.POST("/register", controllers.Register)
	r.Run()
}
