package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	App   AppConfig
	DB    DBConfig
	Redis RedisConfig
	JWT   JWTConfig
	Cors  CorsConfig
}

type DBConfig struct {
	Address      string
	MaxOpenConns int64
	MaxIdleConns int64
	MaxIdleTime  string
}

type RedisConfig struct {
	Address  string
	Password string
	DB       int64
	Protocol int64
}

type JWTConfig struct {
	Secret  string
	Expires int64
}

type AppConfig struct {
	Port    string
	Env     string
	Version string
}

type CorsConfig struct {
	AllowOrigins     string
	AllowMethods     string
	AllowHeaders     string
	ContentLength    string
	MaxAge           int
	AllowCredentials bool
}

var Envs = LoadConfig()

func LoadConfig() *Config {
	viper.AddConfigPath("../")
	viper.SetConfigType("env")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	return &Config{
		App: AppConfig{
			Port:    viper.GetString("PORT"),
			Env:     viper.GetString("ENV"),
			Version: viper.GetString("VERSION"),
		},
		DB: DBConfig{
			Address:      viper.GetString("DB_ADDR"),
			MaxOpenConns: viper.GetInt64("DB_MAX_OPEN_CONNS"),
			MaxIdleConns: viper.GetInt64("DB_MAX_IDLE_CONNS"),
			MaxIdleTime:  viper.GetString("DB_MAX_IDLE_TIME"),
		},
		Redis: RedisConfig{
			Address:  viper.GetString("REDIS_ADDR"),
			Password: viper.GetString("REDIS_PASSWORD"),
			DB:       viper.GetInt64("REDIS_DB"),
			Protocol: viper.GetInt64("REDIS_PROTOCOL"),
		},
		JWT: JWTConfig{
			Secret:  viper.GetString("JWT_SECRET"),
			Expires: viper.GetInt64("JWT_EXPIRATIONS_IN_SECOND"),
		},
		Cors: CorsConfig{
			AllowOrigins:     viper.GetString("ALLOWED_ORIGINS"),
			AllowMethods:     viper.GetString("ALLOWED_METHODS"),
			AllowHeaders:     viper.GetString("ALLOWED_HEADERS"),
			ContentLength:    viper.GetString("CONTENT_LENGTH"),
			MaxAge:           viper.GetInt("MAX_AGE"),
			AllowCredentials: viper.GetBool("ALLOW_CREDENTIALS"),
		},
	}
}
