package routes

import (
	"github.com/MohamedmuhsinJ/shopify/controllers"
	"github.com/MohamedmuhsinJ/shopify/middlewares"
	"github.com/gin-gonic/gin"
)

func UserRoutes(c *gin.Engine) {

	c.POST("/signup", controllers.Signup)
	c.GET("/signup/do", controllers.Do)
	c.POST("/login", controllers.Login)
	c.POST("/login/otp", controllers.SentOtp)
	c.POST("/login/checkOtp", controllers.CheckOtp)
	c.POST("/login/forgetPassword", controllers.ForgetPassword)

	user := c.Group("/")
	user.Use(middlewares.UserAuth())
	{

		user.GET("/home", controllers.UserHome)
		user.GET("/products", controllers.ProductView)
		user.GET("/products/search", controllers.ProductSearch)
		user.GET("/products/sort", controllers.ProductSort)
		user.POST("/addaddress", controllers.Address)
		user.GET("/showaddress", controllers.ShowAddress)
		user.PUT("/editaddress", controllers.EditAddress)
		user.POST("/addtocart", controllers.AddToCart)
		user.GET("/cart", controllers.ViewCart)
		user.GET("/editcart", controllers.EditCart)
		user.GET("/checkout", controllers.CheckOut)

	}
}
