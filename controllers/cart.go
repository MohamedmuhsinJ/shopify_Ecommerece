package controllers

import (
	"strconv"

	"github.com/MohamedmuhsinJ/shopify/database"
	"github.com/MohamedmuhsinJ/shopify/models"
	"github.com/gin-gonic/gin"
)

func AddToCart(c *gin.Context) {
	email := c.GetString("user")
	var user models.User
	var products models.Product
	var productDetails struct {
		Prodct_id uint
		Quantity  uint
	}
	database.Db.Where("email=?", email).Find(&user)
	c.BindJSON(&productDetails)
	database.Db.First(&products, productDetails.Prodct_id)
	totalPrice := products.Price * productDetails.Quantity
	prodID := productDetails.Prodct_id
	proQuantity := productDetails.Quantity
	cart := models.Cart{
		ProductID:  productDetails.Prodct_id,
		Quantity:   productDetails.Quantity,
		UserID:     user.ID,
		TotalPrice: totalPrice,
	}
	var Cart []models.Cart
	//get cart details of user
	database.Db.Where("user_id =?", user.ID).Find(&Cart)
	//checking if the product was already in the cart or not if
	for _, v := range Cart {
		if v.ProductID == prodID {
			total := (proQuantity + v.Quantity) * products.Price
			database.Db.Model(&Cart).Where("user_id=? and product_id=?", user.ID, prodID).Updates(models.Cart{Quantity: proQuantity + v.Quantity, TotalPrice: total})
			c.JSON(200, gin.H{
				"message":  "quantity updated from ",
				"quantity": v.Quantity + +proQuantity,
			})
			c.Abort()
			return
		}
	}
	rec := database.Db.Create(&cart)
	if rec.Error != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"error": rec.Error.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "product added to cart",
	})

}

type CartInfo struct {
	ProductName string
	Price       string
	Image       string
	Stock       uint
	Quantity    string
	TotalPrice  uint
}

func ViewCart(c *gin.Context) {
	var cartinfo []CartInfo
	var user models.User

	userEmail := c.GetString("user")
	database.Db.Where("email=?", userEmail).Find(&user)
	rec := database.Db.Raw("select products.product_name,products.price,products.image,products.stock,carts.quantity,carts.total_price from carts join products on products.product_id=carts.product_id where carts.user_id=?", user.ID).Scan(&cartinfo)

	if rec.Error != nil {
		c.AbortWithStatusJSON(400, gin.H{"err": rec.Error.Error()})
		return
	}
	Total := database.Db.Raw("select sum(total_price) as total from carts where user_id=?", user.ID)
	if Total.Error != nil {
		c.AbortWithStatusJSON(400, gin.H{"err": Total.Error.Error()})
		return
	}
	for i, _ := range cartinfo {
		c.JSON(200, gin.H{
			"prodcuts": cartinfo[i],
		})
	}

}

func EditCart(c *gin.Context) {
	var producsts models.Product
	var user models.User
	var cart models.Cart
	userEmail := c.GetString("user")
	database.Db.Where("email=?", userEmail).Find(&user)
	produc := c.Query("product")
	prodId, _ := strconv.Atoi(produc)
	quant := c.Query("quantity")
	quantity, _ := strconv.Atoi(quant)
	database.Db.Where("product_id=?", prodId).Find(&producsts)
	total := producsts.Price * uint(quantity)
	if quantity >= 1 {
		database.Db.Model(&cart).Where("product_id=? and user_id=?", prodId, user.ID).Updates(models.Cart{Quantity: uint(quantity), TotalPrice: total})
	} else if quantity <= 0 {
		database.Db.Where("user_id=? and product_id=?", prodId, user.ID).Delete(&models.Cart{})
	}
	c.JSON(200, gin.H{
		"message": "go to cart",
	})
}
