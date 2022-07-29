package utils

import "github.com/spf13/viper"

type Config struct {
	DbHost 		string	`mapstructure:"DB_HOST"`
	DbPort 		string	`mapstructure:"DB_PORT"`
	DbName 		string	`mapstructure:"DB_NAME"`
	DbDriver  string	`mapstructure:"DB_DRIVER"`
}

func (c *Config) GetDbSource() string {
	return c.DbDriver + "//" + c.DbHost+":" + c.DbPort
}

func LoadEnv() {
	viper.AutomaticEnv()
	viper.SetConfigName("app")
	viper.SetConfigType("env")
}