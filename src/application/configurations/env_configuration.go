package configurations

import (
	"github.com/spf13/viper"
)

func EnvConfiguration() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
}

func GetEnv(key string) string {
	return viper.GetString(key)
}
