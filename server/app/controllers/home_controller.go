package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/unrolled/render"
)

func Home(c *gin.Context) {
	render := render.New(render.Options{
		Layout: "layout",
	})

	_ = render.HTML(c.Writer, http.StatusOK, "home", map[string]interface{}{
		"title": "Home",
		"body":  "Hello World",
	})
}
