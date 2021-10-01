package middlewares

import (
	"fmt"
	"maker/models"
	"net/http"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

var identityKey = "maker"

// type login struct {
// 	Name     string `form:"name" json:"Name" binding:"required"`
// 	Password string `form:"password" json:"password" binding:"required"`
// }

func NewAuth() (*jwt.GinJWTMiddleware, error) {
	return jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "maker",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			fmt.Println("-------PayloadFunc-----", data.(*models.User))
			if v, ok := data.(*models.User); ok {
				return jwt.MapClaims{
					identityKey: v.Name,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			fmt.Println("-------IdentityHandler-----", claims)
			return &models.User{
				Name: claims[identityKey].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var userParams models.User
			if err := c.ShouldBind(&userParams); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			// 登陆验证
			user := &models.User{
				Name: userParams.Name,
			}
			err := user.FindByName()
			fmt.Println("-------Authenticator-----", user, err, userParams)
			if err == nil && user.Password == userParams.Password {
				return user, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(*models.User); ok {
				err := v.FindByName()
				fmt.Println("-------Authorizator-----", v, err)
				if err == nil {
					return true
				}
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.Redirect(http.StatusFound, "/login")
		},
		LoginResponse: func(c *gin.Context, i int, s string, t time.Time) {
			c.Set("tokenString", s)
			c.Set("expire", t)
			c.Next()
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "cookie: token",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "maker",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})
}
