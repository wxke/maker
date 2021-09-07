package middlewares

import (
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
)

func NewCsrf() gin.HandlerFunc {
	return csrf.Middleware(csrf.Options{
		Secret: "maker",
		ErrorFunc: func(c *gin.Context) {
			c.String(400, "CSRF token mismatch")
			c.Abort()
		},
	})
}
