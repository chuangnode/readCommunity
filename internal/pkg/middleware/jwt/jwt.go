package jwt

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"readCommunity/internal/pkg/errcode"
	"readCommunity/internal/pkg/utils/jwt"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}
		code = errcode.Success
		token := c.Query("token")
		if token == "" {
			code = errcode.ErrTokenInvalid
		} else {
			claims, err := jwt.ParseToken(token)
			if err != nil {
				code = errcode.ErrTokenInvalid
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = errcode.ErrTokenExpired
			}
		}
		if code != errcode.Success {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg": errcode.GetMsg(code),
				"data": data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
