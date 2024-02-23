package controllers

import (
	"html/template"
	"http_gin/src/libs"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HomeController struct{}

func (hc *HomeController) Index(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"main.tmpl.html",
		gin.H{
			"title":        "Página Principal",
			"currentRoute": "home",
			"content": template.HTML(
				libs.Render(
					"src/templates/pages/home/index.tmpl.html",
					nil,
				),
			),
		},
	)
}
