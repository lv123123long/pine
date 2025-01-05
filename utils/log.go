package utils

import (
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
)

func SlogMiddleware(logger *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// 处理请求前
		c.Next()  // 继续处理请求

		// 处理请求后
		cost := time.Since(start)
		logger.Info("handled request",
            slog.String("method", c.Request.Method),
            slog.String("path", path),
            slog.String("query", raw),
            slog.Int("status", c.Writer.Status()),
            slog.Duration("cost", cost),
        )
	}
}