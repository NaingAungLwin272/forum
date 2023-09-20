package config

import (
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	DBUri string `mapstructure:"MONGODB_LOCAL_URI"`
	Port  string `mapstructure:"PORT"`

	GrpcServerAddress string `mapstructure:"GRPC_SERVER_ADDRESS"`

	Origin string `mapstructure:"CLIENT_ORIGIN"`
}

func LoadConfig(path string) (config Config, err error) {
	env := os.Getenv("BUILD_ENV")
	viper.AddConfigPath(path)

	if env == "" {
		viper.SetConfigType("env")
		viper.SetConfigName("app")
	} else {
		viper.SetConfigType("env")
		viper.SetConfigName(env)
	}

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
