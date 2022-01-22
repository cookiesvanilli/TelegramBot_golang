package config

import (
	"github.com/spf13/viper"
)

type Messages struct {
	Errors
	Responses
}

type Errors struct {
	Default      string `mapstructure:"default"`
	InvalidURL   string `mapstructure:"invalid_url"`
	Unauthorized string `mapstructure:"unauthorized"`
	UnableToSave string `mapstructure:"unable_to_save"`
}

type Responses struct {
	Start             string `mapstructure:"start"`
	AlreadyAuthorized string `mapstructure:"already_authorized"`
	SavedSuccessfully string `mapstructure:"saved_successfully"`
	UnknownCommand    string `mapstructure:"unknown_command"`
}

type Config struct {
	Token             string
	PocketConsumerKey string
	AuthServerURL     string
	TelegramBotURL    string `mapstructure:"bot_url"`
	DBPath            string `mapstructure:"db_file"`

	Messages Messages
}

func Init() (*Config, error) {
	if err := setUpViper(); err != nil {
		return nil, err
	}
	var cfg Config

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	if err := viper.UnmarshalKey("messages.responses", &cfg.Messages.Responses); err != nil {
		return nil, err
	}

	if err := viper.UnmarshalKey("messages.errors", &cfg.Messages.Errors); err != nil {
		return nil, err
	}

	if err := parseEnv(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil

}

//парсим переменные окружения
func parseEnv(cfg *Config) error {

	if err := viper.BindEnv("token"); err != nil {
		return err
	}

	cfg.Token = viper.GetString("token")

	if err := viper.BindEnv("consumer_key"); err != nil {
		return err
	}

	cfg.PocketConsumerKey = viper.GetString("consumer_key")

	if err := viper.BindEnv("auth_server_url"); err != nil {
		return err
	}

	cfg.AuthServerURL = viper.GetString("auth_server_url")

	return nil
}

func setUpViper() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("main")

	return viper.ReadInConfig()
}
