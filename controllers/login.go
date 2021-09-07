package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
}

func LoginRegister(router *gin.Engine) {
	loginController := &LoginController{}
	router.LoadHTMLFiles("templates/login/index.tmpl")
	login := router.Group("/login")
	{
		login.GET("/", loginController.index)
	}
}

func (l LoginController) index(c *gin.Context) {
	c.HTML(http.StatusOK, "login/index.tmpl", gin.H{})
}
