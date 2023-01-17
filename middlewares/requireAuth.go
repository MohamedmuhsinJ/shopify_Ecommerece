package middlewares

import (
	"github.com/MohamedmuhsinJ/shopify/controllers"
	"github.com/gin-gonic/gin"
)

func UserAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString, _ := ctx.Cookie("Authorization")
		if tokenString == "" {
			ctx.JSON(401, gin.H{
				"error": "Request does not contain token ",
			})
			ctx.Abort()
			return
		}
		err := controllers.Validate(tokenString)
		ctx.Set("user", controllers.Val)
		if err != nil {
			ctx.JSON(401, gin.H{
				"errorj": err.Error(),
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}

func AdminAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString, _ := ctx.Cookie("AdminAuthorization")
		if tokenString == "" {
			ctx.JSON(401, gin.H{
				"error": "Request does not contain token ",
			})
			ctx.Abort()
			return
		}
		err := controllers.Validate(tokenString)
		ctx.Set("admin", controllers.Val)
		if err != nil {
			ctx.JSON(401, gin.H{
				"error": err.Error(),
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
