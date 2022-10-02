package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/iannealer/go_playground/Ecommerce-api-server/controllers"
)

func UserRoutes(route *gin.Engine) {
	route.POST("/users/signup", controllers.SignUp())
	route.POST("/users/login", controllers.Login())
	route.POST("/admin/addProduct", controllers.ProductViewerAdmin())
	route.GET("/users/productView", controllers.SearchProduct())
	route.GET("/users/search", controllers.SearchProductByQuery())
}
