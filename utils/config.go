package utils

import "github.com/spf13/viper"

type Config struct {
	DbUser		string  `mapstructure:"DB_USER"`
	DbPass		string	`mapstructure:"DB_PASS"`
	DbHost 		string	`mapstructure:"DB_HOST"`
	DbPort 		string	`mapstructure:"DB_PORT"`
	DbName 		string	`mapstructure:"DB_NAME"`
	DbDriver  string	`mapstructure:"DB_DRIVER"`
}

func (c *Config) GetDbSource() string {
	return c.DbDriver + "//" + c.DbUser + ":" + c.DbPass + "@" + c.DbHost+":" + c.DbPort + "/" + c.DbName
}

func LoadEnv() {
	viper.AutomaticEnv()
	viper.SetConfigName("app")
	viper.SetConfigType("env")
}