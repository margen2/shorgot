package config

import (
	"os"

	"github.com/spf13/viper"
)

var (
	StringDBConnection = ""
	Port               = 0
	SecretKey          []byte
)

func Load() {
	viper.SetConfigFile("ENV")
	viper.ReadInConfig()
	viper.AutomaticEnv()

	StringDBConnection = viper.GetString("DB_CONNECTION")
	Port = viper.GetInt("PORT")

	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}
