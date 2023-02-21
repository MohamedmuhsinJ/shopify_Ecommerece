package controllers

import (
	"net/http"

	"github.com/MohamedmuhsinJ/shopify/database"
	"github.com/MohamedmuhsinJ/shopify/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func AdminSignup(c *gin.Context) {
	var admin models.Admin
	if c.ShouldBind(&admin) != nil {
		c.JSON(400, gin.H{
			"error": "failed to read data",
		})
		c.Abort()
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(admin.Password), 10)
	if err != nil {
		c.JSON(502, gin.H{
			"error": "failed to hash password",
		})
		c.Abort()
		return
	}
	admin.Password = string(hash)
	res := database.Db.Create(&admin)
	if res.Error != nil {
		c.JSON(400, gin.H{
			"error": "failed to store data in db",
		})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{
		"message": "admin plese go to login page",
	})
}
func AdminLogin(c *gin.Context) {
	var admin models.Admin
	var body models.Admin
	if c.ShouldBind(&body) != nil {
		c.JSON(400, gin.H{
			"error": "failed to read data",
		})
		c.Abort()
		return
	}
	database.Db.First(&admin, "email=?", body.Email)
	if admin.Email == "" {
		c.JSON(502, gin.H{
			"error": "invalid email plese signup",
		})
		c.Abort()
		return
	}
	err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(body.Password))
	if err != nil {
		c.JSON(400, gin.H{
			"error": "wrong password",
		})
		c.Abort()
		return
	}
	tokenString, err := GenerateToken(body.Email)
	token := tokenString["Token"]
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("AdminAuthorization", token, 3600*24*30, "", "", false, true)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid tokenString",
		})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user":  admin.Email,
		"token": tokenString,
	})
}

func AdminLogout(c *gin.Context) {
	token, _ := c.Get("Token")
	if token == "" {
		c.JSON(401, gin.H{
			"error": "token error",
		})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{
		"message": "admin logedout successfully",
		"token":   token,
	})
}

func AdminDashboard(c *gin.Context) {
	email := c.GetString("admin")
	c.JSON(http.StatusOK, gin.H{
		"Email": email,
	})
}

func UserSearch(c *gin.Context) {
	var user models.User
	name := c.Query("name")
	database.Db.Select("first_name,last_name,email,phone").Where("first_name LIKE ?", name+"%").Find(&user)

	c.JSON(200, gin.H{
		"name":  user.FirstName + " " + user.LastName,
		"email": user.Email,
		"phone": user.Phone,
	})
}
func UserBlock(c *gin.Context) {
	var user models.User
	email := c.Query("email")
	database.Db.Model(&user).Where("email=?", email).Update("block_status", "true")
	c.JSON(200, gin.H{
		"message": user.Email + " " + " blocked",
	})
}

func UserUnblock(c *gin.Context) {
	var user models.User
	email := c.Query("email")
	database.Db.Model(&user).Where("email=?", email).Update("block_status", "false")
	c.JSON(200, gin.H{
		"message": user.FirstName + " " + user.LastName + " successfully unblocked",
	})
}
