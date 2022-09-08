package utils

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

type Configuration struct {
	Server ServerConfig `yaml:"server"`
}

type ServerConfig struct {
	Port string `yaml:"port"`
}

var vp *viper.Viper

var Config Configuration

func NewConfiguration() *Configuration {
	var configuration Configuration
	env := getEnv()

	viper := viper.New()

	viper.AddConfigPath("utils")
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")

	err := viper.ReadInConfig()
	v := viper.Sub(env)

	if err != nil {
		log.Panic(err.Error())
	}

	err = v.Unmarshal(&configuration)

	return &configuration
}

func getEnv() string {

	env := os.Getenv("ACTIVE_PROFILE")

	if env != "" {
		return env
	}

	return "stage"
}
