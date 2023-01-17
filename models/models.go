package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName    string `json:"first_name" validate:"required,min=2,max=100"`
	LastName     string `json:"last_name"  validate:"required,min=2,max=100"`
	Email        string `json:"email" gorm:"unique" validate:"email,required" `
	Password     string `json:"password" validate:"required,min=6"`
	Phone        string `json:"phone"  validate:"required"`
	BlockStatus  bool   `json:blockStatus`
	Token        string `json:"token"`
	RefreshToken string `json:"refereshToken" `
}

type Admin struct {
	Email    string `json:"email" gorm:"unique" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type Product struct {
	ProductId   uint   `json:"product_id" gorm:"primaryKey;not null;autoIncrement"`
	ProductName string `json:"product_name" gorm:"not null"`
	Price       uint   `json:"price" gorm:"not null"`
	ActualPrice uint   `json:"actual_Price" gorm:"not null"`
	Image       string `json:"image" gorm:"not null"`
	SideImage   string `json:"side-image"`
	ZoomImage   string `json:"zoom_image"`
	Stock       uint   `json:"stock"`
	Color       string `json:"color"`
	Description string `json:"description"`
	Discount    uint   `json:"discount"`
	Brand       Brand
	BrandId     uint `json:"brand_id"`
	Cart        Cart
	CartID      uint `json:"cart_id`
	Category    Category
	CategoryID  uint
	ShoeSize    ShoeSize
	ShoeSizeID  uint
	//WishList WishList
	//WishListID uint

}

type Brand struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Brands   string `json:"brands" gorm:"not null"`
	Discount uint   `json:"discount"`
}

type Category struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Category string `json:"category"`
}

type ShoeSize struct {
	ID   uint `json:"id" gorm:"primaryKey"`
	Size uint `json:"size"`
}

type Cart struct {
	CartID     uint `json:"cart_id" gorm:"primaryKey"`
	UserID     uint `json:"user-id"`
	ProductID  uint `json:"product_id"`
	Quantity   uint `json:"quantity"`
	TotalPrice uint `json:"total-price"`
}
