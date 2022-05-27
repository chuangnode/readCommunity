package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	book "readCommunity/internal/bookserver/api/v1"
	"readCommunity/internal/pkg/middleware/cors"
	"readCommunity/internal/pkg/middleware/jwt"
	"readCommunity/internal/pkg/middleware/logger"
	"readCommunity/internal/userserver/api/v1"
)

func NewRouter(r *gin.Engine) {
	r.Use(cors.Cors())
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.POST("/auth", api.GetAuth)
	r.POST("/register", api.Register)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1 := r.Group("/api/v1")
	v1.Use(jwt.JWT())
	{

		v1.POST("/login", api.Login)

		//嵌套路由组
		/*book := v1.Group("book")
		book.POST("/add", func(context *gin.Context) {})*/

		//书籍管理
		//发布书籍
		v1.POST("/book", book.AddBook)
		//修改书籍
		v1.PUT("/book/:id", book.EditBook)
		// 删除书籍
		v1.DELETE("/book/:id", book.DeleteBook)
		// 获取书籍列表
		v1.GET("/book", book.BookList)
		// 获取单书信息
		v1.GET("/book/:id", book.BookInfo)
	}
}
