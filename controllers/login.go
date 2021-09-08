package controllers

import (
	"net/http"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
)

type LoginController struct {
}

func LoginRegister(router *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) {
	loginController := &LoginController{}

	login := router.Group("/login")
	{
		login.GET("/", loginController.index)
		login.POST("/", loginController.auth)
	}
}

func (l LoginController) index(c *gin.Context) {
	c.HTML(http.StatusOK, "login/index.tmpl", gin.H{
		"title": "登录",
		"token": csrf.GetToken(c),
	})
}

func (l LoginController) auth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"title": "登录",
	})
}
