package configurations

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func EnvConfiguration() {
	viper.SetConfigName("configs")
	viper.SetConfigType("env")

	viper.AddConfigPath(".")
	viper.AddConfigPath("/home/bed/Documentos/Code/BackEnd/Go/oohferta/")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
}

func GetEnv(key string) string {
	value := viper.GetString(key)
	os.Setenv(key, value)

	return value
}
