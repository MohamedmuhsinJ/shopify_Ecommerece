package routes

import (
	"github.com/MohamedmuhsinJ/shopify/controllers"
	"github.com/MohamedmuhsinJ/shopify/middlewares"
	"github.com/gin-gonic/gin"
)

func AdminRoutes(c *gin.Engine) {
	c.POST("/adminsignup", controllers.AdminSignup)
	c.POST("/adminlogin", controllers.AdminLogin)
	// c.POST("/login/forgetPssword", contrlers.ForgetPassword)
	admin := c.Group("/admin")
	admin.Use(middlewares.AdminAuth())

	{
		admin.GET("/dashboard", controllers.AdminDashboard)
		admin.GET("/usersearch", controllers.UserSearch)
		admin.PATCH("/userblock", controllers.UserBlock)
	}
}
