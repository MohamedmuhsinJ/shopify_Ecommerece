package controllers

import (
	"github.com/MohamedmuhsinJ/shopify/database"
	"github.com/MohamedmuhsinJ/shopify/models"
	"github.com/gin-gonic/gin"
)

var Products []struct {
	ProductID   uint
	ProductName string
	ActualPrice string
	Price       string
	Image       string
	SideImage   string
	ZoomImage   string
	Description string
	Color       string
	Brands      string
	Stock       uint
	Category    string
	Size        uint
}

func Filter(c *gin.Context) {
	var brandFilter models.Brand
	var categoryFilter models.Category
	var sizeFilter models.ShoeSize
	if brand := c.Query("brandSearch"); brand != "" {
		brandFiles := database.Db.Where("brands=?", brand).Find(&brandFilter)
		if brandFiles.Error != nil {
			c.JSON(400, gin.H{
				"error": brandFiles.Error.Error(),
			})
			c.Abort()
			return
		}
	}
	if category := c.Query("categorySearch"); category != "" {
		categoryFiles := database.Db.Where("category=?", category).Find(&categoryFilter)
		if categoryFiles.Error != nil {
			c.JSON(400, gin.H{
				"error": categoryFiles.Error.Error(),
			})
			c.Abort()
			return
		}
	}

	if size := c.Query("sizeSearch"); size != "" {
		sizeFiles := database.Db.Where("size=?", size).Find(&sizeFilter)
		if sizeFiles.Error != nil {
			c.JSON(400, gin.H{
				"error": sizeFiles.Error.Error(),
			})
			c.Abort()
			return
		}
	}
}
