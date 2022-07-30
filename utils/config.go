package utils

import "github.com/spf13/viper"

type Config struct {
	DbUser		string  `mapstructure:"DB_USER"`
	DbPass		string	`mapstructure:"DB_PASS"`
	DbHost 		string	`mapstructure:"DB_HOST"`
	DbPort 		string	`mapstructure:"DB_PORT"`
	DbName 		string	`mapstructure:"DB_NAME"`
	DbDriver    string	`mapstructure:"DB_DRIVER"`
}

var config *Config

func GetConfig() *Config {
	if config == nil {
		config = &Config{}
	}
	return config
}

func (c *Config) GetDbSource() string {
	return  c.DbUser + ":" + c.DbPass + "@tcp(" + c.DbHost+":" + c.DbPort + ")/" + c.DbName + "?charset=utf8&parseTime=True&loc=Local"
}

func LoadEnv() {
	viper.SetConfigType("env")
	viper.SetConfigName("app")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(GetConfig())
}