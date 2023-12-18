package app

import (
	"mini-marketplace/app/controllers"

	"github.com/gin-gonic/gin"
)

func (s *Server) InitializeRoutes() {
	s.Router = gin.Default()
	s.Router.GET("/", controllers.Home)
}
