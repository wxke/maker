package routers

import (
	"log"
	"maker/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	authMiddleware, err := middlewares.NewAuth()
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	router.GET("/index", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{
				"code": http.StatusOK,
			},
		)
	})
	router.POST("/login", authMiddleware.LoginHandler)

	return router
}
