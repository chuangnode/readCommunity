package cors

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
	"net/http"
)

type corsWrapper struct {
	*cors.Cors
}

//Cors add cors header
func (c corsWrapper)build() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		c.HandlerFunc(ctx.Writer, ctx.Request)
		if ctx.Request.Method == http.MethodOptions{
			ctx.AbortWithStatus(http.StatusOK)
		}
		ctx.Next()
	}
}

func Cors() gin.HandlerFunc  {
	return corsWrapper{cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"OPTIONS","GET","POST","PUT","PATCH","DELETE"},
		AllowedHeaders: []string{"Origin","Content-Type","Accept","Authorization"},
		ExposedHeaders: []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge: 12 * 3600,
	})}.build()
}
