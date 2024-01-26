package routes

import (
	"mini-marketplace/handlers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/signup", handlers.SignUp())
	incomingRoutes.POST("/users/login", handlers.Login())
	incomingRoutes.POST("/admin/add-product", handlers.ProductViewerAdmin())
	incomingRoutes.GET("/users/product-view", handlers.SearchProduct())
	incomingRoutes.GET("users/search", handlers.SearchProductByQuery())
}
