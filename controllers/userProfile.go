package controllers

import (
	"strconv"

	"github.com/MohamedmuhsinJ/shopify/database"
	"github.com/MohamedmuhsinJ/shopify/models"
	"github.com/gin-gonic/gin"
)

type Profile struct {
	Name     string `json:"name,omitempty"`
	Phone    uint   `json:"phone,omitempty"`
	Pincode  uint   `json:"pincode,omitempty"`
	Area     string `json:"area,omitempty"`
	Landmark string `json:"landmark,omitempty"`
	City     string `json:"city,omitempty"`
}

func Address(c *gin.Context) {
	useremail := c.GetString("user")
	var user models.User
	database.Db.Raw("select id,first_name,last_name,phone from users where email =?", useremail).Scan(&user)
	area := c.PostForm("area")
	phone, _ := strconv.Atoi(user.Phone)

	landmark := c.PostForm("landmark")
	city := c.PostForm("city")
	pin := c.PostForm("pincode")
	pincode, _ := strconv.Atoi(pin)
	address := models.Address{
		UserId:      user.ID,
		Name:        user.FirstName + " " + user.LastName,
		PhoneNumber: phone,
		Email:       useremail,
		Area:        area,
		Landmark:    landmark,
		City:        city,
		Pincode:     pincode,
	}

	rec := database.Db.Create(&address)
	if rec.Error != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"error": rec.Error.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"address": address,
	})
}

func ShowAddress(c *gin.Context) {
	var user models.Address
	email := c.GetString("user")
	rec := database.Db.Where("email =?", email).Find(&user)
	if rec.Error != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"error": rec.Error.Error(),
		})
		return

	}
	c.JSON(200, gin.H{
		"address": user,
	})
}

func EditAddress(c *gin.Context) {
	var profile Profile
	var user models.Address
	email := c.GetString("user")
	if err := c.ShouldBindJSON(&profile); err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	rec := database.Db.Model(user).Where("email=?", email).Updates(models.Address{Name: profile.Name, Pincode: int(profile.Pincode), Landmark: profile.Landmark, Area: profile.Area, City: profile.City, PhoneNumber: int(profile.Phone)})
	if rec.Error != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"err": rec.Error.Error(),
		})
	}
	c.JSON(200, gin.H{
		"msg": "address updated",
	})

}
