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
		login.POST("/", authMiddleware.LoginHandler, loginController.auth)
	}
}

func (l LoginController) index(c *gin.Context) {
	c.HTML(http.StatusOK, "login/index.tmpl", gin.H{
		"title": "登录",
		"token": csrf.GetToken(c),
	})
}

func (l LoginController) auth(c *gin.Context) {
	tokenString, _ := c.Get("tokenString")
	expire, _ := c.Get("expire")
	c.SetCookie("token", tokenString.(string), 3600, "/", "127.0.0.1", false, false)
	c.JSON(http.StatusOK, gin.H{
		"tokenString": tokenString,
		"expire":      expire,
	})
}
