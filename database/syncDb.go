package database

import "github.com/MohamedmuhsinJ/shopify/models"

func SyncDb() {

	Db.AutoMigrate(
		&models.User{},
		&models.Admin{},
		&models.Brand{},
		&models.Cart{},
		&models.Category{},
		&models.ShoeSize{},
		&models.Product{},
		&models.Address{},
		&models.OrderedItem{},
		&models.Orders{},
	)
}
