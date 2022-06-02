package config

type Config struct {
	// 数据库配置
	DB struct {
		Host     string `json:"host" mapstructure:"host"`
		Port     string `json:"port" mapstructure:"port"`
		User     string `json:"user" mapstructure:"user"`
		Password string `json:"password" mapstructure:"password"`
		Name     string `json:"name" mapstructure:"name"`
	} `json:"db" mapstructure:"db"`
}
