package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lv123123long/pine/utils"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "token is none"})
			ctx.Abort()
			return
		}
		username, err := utils.ParseJWT(token)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "faewf"})
			ctx.Abort()
			return
		}

		ctx.Set("usrname", username)
		ctx.Next()
	}
}
