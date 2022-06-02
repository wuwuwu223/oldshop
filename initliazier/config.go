package initliazier

import (
	"fmt"
	"github.com/spf13/viper"
	"oldshop/global"
)

//初始化配置文件
func InitConfig() {
	v := viper.New()
	v.SetConfigType("json")
	configFileName := fmt.Sprintf("config.json")
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := v.Unmarshal(&global.Config); err != nil {
		panic(err)
	}
}
