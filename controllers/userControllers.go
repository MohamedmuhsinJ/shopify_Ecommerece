package controllers

import (
	"net/http"

	"github.com/MohamedmuhsinJ/shopify/database"
	"github.com/MohamedmuhsinJ/shopify/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	var body struct {
		FirstName string
		LastName  string
		Email     string
		Password  string
		Phone     string
	}
	if c.ShouldBind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read body",
		})
		c.Abort()
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "failed to hash",
		})
		c.Abort()
		return
	}
	user := models.User{FirstName: body.FirstName, LastName: body.LastName, Email: body.Email, Password: string(hash), Phone: body.Phone}
	result := database.Db.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to create user",
		})
		c.Abort()
	}
	email := c.Query(user.Email)
	c.String(http.StatusAccepted, "hello %s", email)
	c.JSON(http.StatusOK, gin.H{
		"messagge": "plese login",
	})
}

func Login(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}
	if c.ShouldBind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to get request",
		})
		c.Abort()
		return
	}
	var user models.User
	//result := database.Db.Where("email=?", body.Email).First(&user)
	database.Db.First(&user, "email=?", body.Email)
	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid email",
		})
		c.Abort()
		return
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "wrong password",
		})
		c.Abort()
		return
	}

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid tokenString",
		})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user": user.FirstName + user.LastName,
	})
}

func UserHome(c *gin.Context) {
	var user models.User
	email := c.GetString("user")

	database.Db.First(&user, "email=?", email)

	c.JSON(200, gin.H{
		"username": user.FirstName + user.LastName,
	})
}

// func lValidate(c *gin.Context) {
// 	User, err := c.Get("user")
// 	if !err {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": "lfailed to create token",
// 		})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": User,
// 	})
// }

func ForgetPassword(c *gin.Context) {
	var user models.User
	var body struct {
		Email string
	}
	if c.ShouldBind(&body) != nil {
		c.JSON(400, gin.H{
			"error": "failed to get request",
		})
		return
	}

	database.Db.First(&user, "email=?", body.Email)
	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid email",
		})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"user": user,
	})
}

func Do(c *gin.Context) {
	email := c.Query("email")
	c.JSON(http.StatusAccepted, gin.H{
		"message": email,
	})
}
