package main

import (
	"github.com/gin-gonic/gin"
	"github.com/iannealer/go_playground/Ecommerce-api-server/routes"
	"log"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))

	r := gin.New()
	r.Use(gin.Logger())
	routes.UserRoutes(r)
	r.Use(middleware.Authentication())

	r.GET("/addtocart", app.AddToCart())
	r.GET("removeitem", app.RemoveItem())
	r.GET("/cartcheckout", app.BuyFromCart())
	r.GET("/instantbuy", app.InstantBuy())

	log.Fatal(r.Run(":" + port))
}
