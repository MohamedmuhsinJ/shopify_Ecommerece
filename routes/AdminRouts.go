package routes

import (
	"github.com/MohamedmuhsinJ/shopify/controllers"
	"github.com/MohamedmuhsinJ/shopify/middlewares"
	"github.com/gin-gonic/gin"
)

func AdminRoutes(c *gin.Engine) {
	c.POST("/adminsignup", controllers.AdminSignup)
	c.POST("/adminlogin", controllers.AdminLogin)
	admin := c.Group("/admin")
	admin.Use(middlewares.AdminAuth())

	{
		admin.GET("/dashboard", controllers.AdminDashboard)
		admin.GET("/logout", controllers.AdminLogout)
		admin.GET("/usersearch", controllers.UserSearch)
		admin.PUT("/userblock", controllers.UserBlock)
		admin.PUT("/userunblock", controllers.UserUnblock)
		admin.GET("/listall", controllers.ListALL)

		admin.POST("/addproduct", controllers.AddProduct)
		admin.PUT("/editproduct/:id", controllers.EditProduct)
		admin.DELETE("/deleteproduct/:id", controllers.DeleteProducts)
		admin.GET("/getproduct/:id", controllers.GetProduct)
	}
}
