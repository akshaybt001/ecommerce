package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type Config struct {
	DBHost           string `mapstructure:"DB_HOST"`
	DBName           string `mapstructure:"DB_NAME"`
	DBUser           string `mapstructure:"DB_USER"`
	DBPassword       string `mapstructure:"DB_PASSWORD"`
	DBSslmode        string `mapstructure:"DB_SSLMODE"`
	TWILIOACCOUNTSID string `mapstructure:"TWILIO_ACCOUNT_SID"`
	TWILIOAUTHTOKEN  string `mapstructure:"TWILIO_AUTHTOKEN"`
	TWILIOSERVICESID string `mapstructure:"TWILIO_SERVICES_ID"`
}

var envs = []string{
	"DB_HOST",
	"DB_NAME",
	"DB_USER",
	"DB_PASSWORD",
	"DB_SSLMODE",
	"TWILIO_ACCOUNT_SID",
	"TWILIO_AUTHTOKEN",
	"TWILIO_SERVICES_ID",
}

func LoadConfig() (Config, error) {
	var config Config

	viper.AddConfigPath("../.env")
	viper.SetConfigFile("../.env")
	viper.ReadInConfig()

	for _, env := range envs {
		if err := viper.BindEnv(env); err != nil {
			return config, err
		}
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	if err := validator.New().Struct(&config); err != nil {
		return config, err
	}

	return config, nil
}
