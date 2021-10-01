package controllers

import (
	"net/http"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type HomeController struct {
}

func HomeRegister(router *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) {
	homeController := &HomeController{}
	home := router.Group("/")
	{
		home.GET("/", authMiddleware.MiddlewareFunc(), homeController.index)
	}
}

func (h HomeController) index(c *gin.Context) {
	c.HTML(http.StatusOK, "home/index.tmpl", gin.H{
		"title": "主页",
	})
}
