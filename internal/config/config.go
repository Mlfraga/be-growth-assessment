package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
    DBUsername string
    DBPassword string
    DBName     string
    DBHost     string
    DBPort     string
    RedisHost  string
    RedisPort  string
    JWTSecret  string
}

func LoadConfig() *Config {
    viper.AddConfigPath(".")
    viper.SetConfigName(".env")
    viper.SetConfigType("env")
    viper.AutomaticEnv()

    err := viper.ReadInConfig()
    if err != nil {
        log.Fatalf("Error while reading config file %s", err)
    }

    var config Config
    err = viper.Unmarshal(&config)
    if err != nil {
        log.Fatalf("Unable to decode into struct %s", err)
    }

    return &config
}
