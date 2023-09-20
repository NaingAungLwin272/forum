package config

import (
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Port           string `mapstructure:"PORT"`
	UserSvcUrl     string `mapstructure:"USER_SVC_URL"`
	FeaturesSvcUrl string `mapstructure:"FEATURES_SVC_URL"`
	CategorySvcUrl string `mapstructure:"CATEGORY_SVC_URL"`
	NotiSvcUrl     string `mapstructure:"NOTI_SVC_URL"`
	BadgeSvcUrl    string `mapstructure:"BADGE_SVC_URL"`
	MailSvcUrl     string `mapstructure:"MAIL_SVC_URL"`
	FRONT_END_URL  string `mapstructure:"FRONT_END_URL"`
	BOT_URL        string `mapstructure:"BOT_URL"`
}

func LoadConfig() (c Config, err error) {
	env := os.Getenv("BUILD_ENV")
	viper.AddConfigPath("./pkg/config/envs")

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

	err = viper.Unmarshal(&c)

	return
}
