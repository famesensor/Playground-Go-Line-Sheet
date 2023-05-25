package configs

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	AppPort                 string `mapstructure:"APP_PORT"`
	LineChannelSecret       string `mapstructure:"LINE_CHANNEL_SECRET"`
	LineChannelToken        string `mapstructure:"LINE_CHANNEL_TOKEN"`
	SheetCredentialFileName string `mapstructure:"SHEET_CREDENTIAL_FILE_NAME"`
	SheetSpreadSheetId      string `mapstructure:"SHEET_SPREAD_SHEET_ID"`
}

var config Config

func InitViper() {
	// TODO: setup switch env
	viper.AddConfigPath(".")
	viper.SetConfigType("env")
	viper.SetConfigName(".env")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
	err = viper.Unmarshal(&config)
	if err != nil {
		fmt.Println(err)
	}
}

func GetConfig() *Config {
	return &config
}
