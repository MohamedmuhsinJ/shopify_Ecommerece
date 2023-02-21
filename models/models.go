package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName   string `json:"first_name" validate:"required,min=2,max=100"`
	LastName    string `json:"last_name"  validate:"required,min=2,max=100"`
	Email       string `json:"email" gorm:"unique" validate:"email,required" `
	Password    string `json:"password" validate:"required,min=6"`
	Phone       string `json:"phone"  validate:"required"`
	BlockStatus bool   `json:"blockStatus"`
	Address     Address
	AddressId   uint `json:"address_id"`
	Cart        Cart
	CartId      uint `json:"cart_id"`
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
	CartID      uint `json:"cart_id"`
	Category    Category
	CategoryID  uint
	ShoeSize    ShoeSize
	ShoeSizeID  uint
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

type Address struct {
	AddressId   uint   `json:"Address_id" gorm:"primaryKey"`
	UserId      uint   `json:"user_id"`
	Name        string `json:"name"`
	PhoneNumber int    `json:"phone_number"`
	Email       string `json:"email"`
	Area        string `json:"area"`
	Landmark    string `json:"landmark"`
	City        string `json:"city"`
	Pincode     int    `json:"pincode"`
}

type Cart struct {
	CartID     uint `json:"cart_id" gorm:"primaryKey"`
	UserID     uint `json:"user_id"`
	ProductID  uint `json:"product_id"`
	Quantity   uint `json:"quantity"`
	TotalPrice uint `json:"total_price"`
}

type OrderedItem struct {
	ID            uint `gorm:"primaryKey"`
	UserID        uint
	ProductID     uint
	OrederID      string
	ProductName   string
	Quantity      uint
	Price         uint
	OrderStatus   string
	PaymentStatus string
	PaymentMethod string
	TotalPrice    uint
}

type Orders struct {
	gorm.Model
	UserID         uint
	OrderId        string `json:"order_id"  gorm:"not null" `
	TotalAmount    uint   `json:"total_amount"  gorm:"not null" `
	PaymentMethod  string `json:"paymentmethod"  gorm:"not null" `
	Payment_Status string `json:"payment_status"   `
	Order_Status   string `json:"order_status"   `
	Address        Address
	Address_id     uint `json:"address_id"  `
}
