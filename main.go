package main

import (
	"os"

	"github.com/MohamedmuhsinJ/shopify/database"
	"github.com/MohamedmuhsinJ/shopify/initalizers"
	routes "github.com/MohamedmuhsinJ/shopify/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	initalizers.LoadEnvVariables()
	database.ConnectToDb()
	database.SyncDb()
}

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	router := gin.Default()
	routes.AdminRoutes(router)
	routes.UserRoutes(router)
	router.Run()

}
