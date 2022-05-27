package initconfig

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"readCommunity/global"
)

//viper读取配置信息，gorm连接mysql
func InitMySQL() {
	dbUser := viper.GetString("Database.mysql.username")
	dbPwd := viper.GetString("Database.mysql.password")
	dbHost := viper.GetString("Database.mysql.host")
	dbPort := viper.GetString("Database.mysql.port")
	dbDatabase := viper.GetString("Database.mysql.database")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser,
		dbPwd, dbHost, dbPort, dbDatabase)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("mysql connect failed, err: %v", err)
	}

	global.DB = db
}
