package config

import (
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Port              string `mapstructure:"PORT"`
	GrpcServerAddress string `mapstructure:"GRPC_SERVER_ADDRESS"`
	Origin            string `mapstructure:"CLIENT_ORIGIN"`
	SENDGRID_API_KEY  string `mapstructure:"SENDGRID_API_KEY"`
	SEND_FROM_NAME    string `mapstructure:"SEND_FROM_NAME"`
	SEND_FROM_ADDRESS string `mapstructure:"SEND_FROM_ADDRESS"`
	SEND_TO_NAME      string `mapstructure:"SEND_TO_NAME"`
	SEND_TO_ADDRESS   string `mapstructure:"SEND_TO_ADDRESS"`
	FRONT_END_URL     string `mapstructure:"FRONT_END_URL"`
	FRONT_ADMIN_URL   string `mapstructure:"FRONT_ADMIN_URL"`
}

func LoadConfig(path string) (config Config, err error) {
	env := os.Getenv("BUILD_ENV")
	viper.AddConfigPath(path)

	if env == "" {
		viper.SetConfigType("env")
		viper.SetConfigName("app")
	} else {
		viper.SetConfigType("env")
		viper.SetConfigName("dev")
	}

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
