package configs

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

const (
	AppEnvLocal      = "local"
	AppEnvDev        = "dev"
	AppEnvTest       = "test"
	AppEnvProduction = "production"
)

var (
	appEnv = os.Getenv("APP_ENV")
)

func InitConfig() {
	viper.BindEnv("APP_NAME")
	viper.BindEnv("APP_ENV")
	viper.BindEnv("APP_PATH")
	viper.BindEnv("APP_LOG_PATH")
	viper.BindEnv("APP_SETTINGS_PATH")

	viper.SetDefault("APP_NAME", "hello")
	viper.SetDefault("APP_ENV", "dev")

	var fileName string
	switch appEnv {
	case AppEnvLocal:
		fileName = "local.yaml"
	case AppEnvDev:
		fileName = "dev.yaml"
	case AppEnvTest:
		fileName = "test.yaml"
	case AppEnvProduction:
		fileName = "production.yaml"
	default:
		fileName = "local.yaml"
	}

	viper.SetConfigType("yaml")

	AppConfigFile := filepath.Join(viper.GetString("APP_SETTINGS_PATH"), "settings", fileName)

	// for develop env
	if _, err := os.Stat(AppConfigFile); os.IsNotExist(err) {
		defaultAppSettingsPath := filepath.Join("configs", viper.GetString("APP_NAME") )
		viper.SetDefault("APP_SETTINGS_PATH", defaultAppSettingsPath)
	}

	AppConfigFile = filepath.Join(viper.GetString("APP_SETTINGS_PATH"), "settings", fileName)

	viper.SetDefault("APP_CONFIG_FILE", AppConfigFile)
	viper.SetConfigFile(viper.GetString("APP_CONFIG_FILE"))
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error settings file: %s \n", err))
	} else {
		fmt.Printf("APP Configs: %+v\n", viper.AllSettings())
	}
}
