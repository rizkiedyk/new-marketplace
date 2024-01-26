package main

import (
	"log"
	"mini-marketplace/config/database"
	"mini-marketplace/handlers"
	"mini-marketplace/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	cli := database.DBSet()

	prodctColl := cli.Database("New-Marketplace").Collection("Products")
	usersColl := cli.Database("New-Marketplace").Collection("Users")

	app := handlers.NewApplication(prodctColl, usersColl)
	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	// router.Use(middleware.Authentication())

	router.GET("/addtocart", app.AddToCart())
	router.GET("/removeitem", app.RemoveItem())
	router.GET("/cart-checkout", app.BuyFromCart())
	router.GET("/instant-buy", app.InstantBuy())

	log.Fatal(router.Run(":" + port))
}
