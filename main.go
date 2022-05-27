package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	_ "readCommunity/initconfig"
	"readCommunity/internal/pkg/middleware/logger"
	"readCommunity/routers"
)

// @swagger:                   2.0
// @title                      Swagger API
// @version                    1.0
// @description                this is a read community
// license.name MIT
// @host                       localhost:8080
// @BasePath                   /api/v1
// @securityDefinitions.basic  BasicAuth
//TODO {{swagger ui访问localhost:8080/swagger/index.html,报错}}
func main() {
	//TODO
	/*
		1.注册、登陆
		2.图书管理、分类管理
		3.收藏、评论、评分
	*/

	// init logger
	if err := logger.InitLogger(); err != nil {
		fmt.Printf("init logger failed, err: %v\n", err)
		return
	}

	r := gin.New()
	gin.SetMode(viper.GetString("GinMode"))
	routers.NewRouter(r)
	r.Run()
}
