package routers

import (
	"log"
	"maker/controllers"
	"maker/middlewares"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Static("/statics", "./statics")
	router.LoadHTMLGlob("templates/**/*")
	// middlewares
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	authMiddleware, err := middlewares.NewAuth()
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
	store := cookie.NewStore([]byte("maker"))
	router.Use(sessions.Sessions("maker_session", store))
	router.Use(middlewares.NewCsrf())

	// routers
	controllers.LoginRegister(router, authMiddleware)

	router.GET("/index", authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"tokenString": "tokenString",
			"expire":      "expire",
		})
	})
	// router.GET("/login")
	// router.POST("/login", authMiddleware.LoginHandler)

	return router
}
