package initialize

import (
	"fmt"
	"github.com/spf13/viper"
)

func LoadConfig() {
	viper.AddConfigPath("configs")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("viper read config failed, ", err)
	}
}
