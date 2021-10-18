package controllers

import (
	"maker/models"
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
		login.GET("/new", loginController.signUp)
		login.POST("/new", loginController.create)
	}
}

func (l LoginController) index(c *gin.Context) {
	c.HTML(http.StatusOK, "login/index.html", gin.H{
		"title": "登录",
		"token": csrf.GetToken(c),
	})
}

func (l LoginController) auth(c *gin.Context) {
	tokenString, _ := c.Get("tokenString")
	// expire, _ := c.Get("expire")
	c.SetCookie("token", tokenString.(string), 3600, "/", "127.0.0.1", false, false)

	c.Redirect(http.StatusFound, "/")
}

func (l LoginController) signUp(c *gin.Context) {
	c.HTML(http.StatusOK, "login/new.html", gin.H{
		"title": "注册",
		"token": csrf.GetToken(c),
	})
}

func (l LoginController) create(c *gin.Context) {
	var userParams models.User
	if err := c.ShouldBind(&userParams); err != nil {
		c.Redirect(http.StatusFound, c.GetHeader("Referer"))
	}
	user := &models.User{
		Name:     userParams.Name,
		Password: userParams.Password,
		Email:    userParams.Email,
	}

	if err := user.Create(); err != nil {
		c.Redirect(http.StatusFound, c.GetHeader("Referer"))
	} else {
		c.Redirect(http.StatusFound, "/login")
	}
}
