package config

import (
	"github.com/spf13/viper"
	"sync"
)

type Config struct {
	viper *viper.Viper
}

var configInstance *Config // package private singleton instance of the configuration
var singleton sync.Once    // package private singleton helper utility

func GetConfig(env string,vendor string) (*viper.Viper,error) {
	// create an instance if not availabe
	singleton.Do(func() {
		configInstance = &Config{viper.New()}
	})

	err := LoadConfig(env,vendor)
	return configInstance.viper,err
}

func LoadConfig(env string,vendor string) (error) {

	// init default config
	config_file_name := vendor + "_" + env+".json"
	configInstance.viper.SetConfigName(config_file_name)
	configInstance.viper.AddConfigPath("./config")
	configInstance.viper.SetConfigType("json")
	configInstance.viper.AutomaticEnv()

	// Find and read the config file
	err := configInstance.viper.ReadInConfig()
	if err != nil {
		println("error")
	}
	return err
}