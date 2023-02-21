package controllers

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/MohamedmuhsinJ/shopify/database"
	"github.com/MohamedmuhsinJ/shopify/models"
	"github.com/gin-gonic/gin"
)

type OrderDetails struct {
	UserID      uint
	ProductID   uint
	ProductName string
	Price       uint
	Quantity    uint
	Total       uint
}

func CheckOut(c *gin.Context) {
	var user models.User
	var orderDetails []OrderDetails
	var address models.Address
	var orderedItems models.OrderedItem
	var cart models.Cart
	var Msg string
	email := c.GetString("user")
	addr := c.Query("addressId")
	add, _ := strconv.Atoi(addr)
	paymentMethod := c.Query("paymentMethod")
	database.Db.Where("email=?", email).Find(&user)
	carts := database.Db.Raw("select products.product_id,products.product_name,products.price,carts.quantity,carts.user_id,total_price from carts join products on products.product_id=carts.product_id where carts.user_id=?", user.ID).Scan(&orderDetails)
	if carts.Error != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"error": carts.Error.Error(),
		})
		return
	}
	var totatCartValue uint
	database.Db.Raw("select sum(total_price) as total from carts where user_id=?", user.ID).Scan(&totatCartValue)
	if paymentMethod == "stripe" {
		Msg = Stripe(totatCartValue)
		if Msg == "" {
			c.AbortWithStatusJSON(400, gin.H{
				"err": "Stripe error",
			})
			return
		} else {
			fmt.Fprint(c.Writer, Msg)
		}
	}
	if paymentMethod == "COD" || paymentMethod == "stripe" {
		for _, v := range orderDetails {
			iD := v.UserID
			pId := v.ProductID
			pName := v.ProductName
			price := v.Price
			quan := v.Quantity
			total := v.Quantity * v.Price
			if paymentMethod == "COD" {
				ordredItems := models.OrderedItem{UserID: iD, ProductID: pId, ProductName: pName, Quantity: quan, Price: price, OrderStatus: "confirmed", PaymentStatus: "pending", PaymentMethod: "COD", TotalPrice: total}
				Msg = "Cash On Delivery order placed"
				rec := database.Db.Create(&ordredItems)
				if rec.Error != nil {
					c.AbortWithStatusJSON(400, gin.H{
						"error": rec.Error.Error(),
					})
					return
				}
			} else if paymentMethod == "stripe" {
				ordredItems := models.OrderedItem{UserID: iD, ProductID: pId, ProductName: pName, Quantity: quan, Price: price, OrderStatus: "confirmed", PaymentStatus: "pending", PaymentMethod: "stripe", TotalPrice: total}

				rec := database.Db.Create(&ordredItems)
				if rec.Error != nil {
					c.AbortWithStatusJSON(400, gin.H{
						"error": rec.Error.Error(),
					})
					return
				}
			}
		}
		c.JSON(200, gin.H{
			"msg": "ordrer placing!!!",
		})

	}

	rec := database.Db.Where("user_id=? and address_id=?", user.ID, add).Find(&address)
	if rec.Error != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"err": rec.Error.Error(),
		})
		return
	}
	c.JSON(300, gin.H{"address": address, "total value": totatCartValue})

	OrderID := OID()
	if paymentMethod == "COD" && add == int(address.AddressId) {
		order := models.Orders{
			UserID:         user.ID,
			Address_id:     uint(add),
			PaymentMethod:  "COD",
			Payment_Status: "Cash On Delivery",
			TotalAmount:    totatCartValue,
			OrderId:        OrderID,
			Order_Status:   "order placed",
		}
		res := database.Db.Create(&order)
		if res.Error != nil {
			c.AbortWithStatusJSON(400, gin.H{"err": res.Error.Error()})
			return
		}
		resc := database.Db.Raw("update ordered_items set oreder_id=?, order_status =?,payment_status=? Where user_id=?", OrderID, "order placed", "COD", user.ID).Scan(&orderedItems)
		if resc.Error != nil {
			c.AbortWithStatusJSON(400, gin.H{"err": resc.Error.Error()})
			return
		}
		rec := database.Db.Raw("delete from carts where user_id=?", user.ID).Scan(&cart)
		if rec.Error != nil {
			c.AbortWithStatusJSON(400, gin.H{"err": resc.Error.Error()})
			return
		}
	} else if paymentMethod == "stripe" && add == int(address.AddressId) {
		order := models.Orders{
			UserID:         user.ID,
			Address_id:     uint(add),
			PaymentMethod:  "Stripe",
			Payment_Status: "Online",
			TotalAmount:    totatCartValue,
			OrderId:        OrderID,
			Order_Status:   "order placed",
		}
		res := database.Db.Create(&order)
		if res.Error != nil {
			c.AbortWithStatusJSON(400, gin.H{"err": res.Error.Error()})
			return
		}
		resc := database.Db.Raw("update ordered_items set oreder_id=?, order_status =?,payment_status=? Where user_id=?", OrderID, "order placed", "Stripe", user.ID).Scan(&orderedItems)
		if resc.Error != nil {
			c.AbortWithStatusJSON(400, gin.H{"err": resc.Error.Error()})
			return
		}
		rec := database.Db.Raw("delete from carts where user_id=?", user.ID).Scan(&cart)
		if rec.Error != nil {
			c.AbortWithStatusJSON(400, gin.H{"err": resc.Error.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "Thank you for Shopping with us!"})
	}

}

func OID() string {

	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(9999999999-1000000000) + 1000000000
	ID := strconv.Itoa(value)
	return "OID" + ID

}
