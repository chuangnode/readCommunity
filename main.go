package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func main()  {
//TODO
	/*
		1.注册、登陆
		2.图书管理、分类管理
		3.收藏、评论、评分
	*/
	viper.AddConfigPath("configs")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("viper read config failed, " , err)
	}
	fmt.Println(viper.GetString("Database.mysql.username"))
}