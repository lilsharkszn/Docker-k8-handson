package config

import (
    "github.com/spf13/viper"
)

// Config stores all configuration of the application.
type Config struct {
    ServerPort         string `mapstructure:"PORT"`
    MongoURI           string `mapstructure:"MONGO_URI"`
    DBName             string `mapstructure:"DB_NAME"`
    JWTSecretKey       string `mapstructure:"JWT_SECRET_KEY"`
    JWTExpirationHours int    `mapstructure:"JWT_EXPIRATION_HOURS"`
    EnableCache        bool   `mapstructure:"ENABLE_CACHE"`
    RedisAddr          string `mapstructure:"REDIS_ADDR"`
    RedisPassword      string `mapstructure:"REDIS_PASSWORD"`
    LogLevel           string `mapstructure:"LOG_LEVEL"`
    LogFormat          string `mapstructure:"LOG_FORMAT"`
}

func LoadConfig(path string) (config Config, err error) {

    viper.AutomaticEnv()
    viper.BindEnv("MONGO_URI")
    viper.BindEnv("PORT")
    viper.BindEnv("DB_NAME")
    viper.SetDefault("PORT", "8080")
    viper.SetDefault("ENABLE_CACHE", false)
    viper.SetDefault("JWT_EXPIRATION_HOURS", 72)

    err = viper.Unmarshal(&config)
    return
}
