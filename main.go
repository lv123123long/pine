package main

import (
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/lv123123long/pine/controllers"
	"github.com/lv123123long/pine/utils"
)

func main() {
	r := gin.Default()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))

	// 添加自定义日志中间件
	r.Use(utils.SlogMiddleware(logger))
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// 注册路由并使用处理器
	greetHandler := controllers.NewGreetHandler(logger)
	r.POST("/login", greetHandler.Login)
	r.POST("/register", greetHandler.Register)
	r.Run()
}
