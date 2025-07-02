package config

import (
	"github.com/spf13/viper"
)

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
	Duration int64
}

type JWTConfig struct {
	Secret         string
	Expires        int64
	RefreshExpires int64
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

type MailConfig struct {
	SMTPHost     string
	SMTPPort     int
	SenderName   string
	SenderEmail  string
	SMTPPassword string
}

// var Envs = LoadConfig()
var App *AppConfig
var DB *DBConfig
var Redis *RedisConfig
var JWT *JWTConfig
var Cors *CorsConfig
var Mail *MailConfig

func LoadConfig() error {
	viper.AddConfigPath("../")
	viper.SetConfigType("env")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	App = &AppConfig{}
	App.Port = viper.GetString("PORT")
	App.Env = viper.GetString("ENV")
	App.Version = viper.GetString("VERSION")

	DB = &DBConfig{}
	DB.Address = viper.GetString("DB_ADDR")
	DB.MaxOpenConns = viper.GetInt64("DB_MAX_OPEN_CONNS")
	DB.MaxIdleConns = viper.GetInt64("DB_MAX_IDLE_CONNS")
	DB.MaxIdleTime = viper.GetString("DB_MAX_IDLE_TIME")

	Redis = &RedisConfig{}
	Redis.Address = viper.GetString("REDIS_ADDR")
	Redis.Password = viper.GetString("REDIS_PASSWORD")
	Redis.DB = viper.GetInt64("REDIS_DB")
	Redis.Protocol = viper.GetInt64("REDIS_PROTOCOL")
	Redis.Duration = viper.GetInt64("REDIS_DURATION")

	JWT = &JWTConfig{}
	JWT.Secret = viper.GetString("JWT_SECRET")
	JWT.Expires = viper.GetInt64("JWT_EXPIRATIONS_IN_SECOND")
	JWT.RefreshExpires = viper.GetInt64("REFRESH_TOKEN_EXPIRATIONS_IN_SECOND")

	Cors = &CorsConfig{}
	Cors.AllowOrigins = viper.GetString("ALLOWED_ORIGINS")
	Cors.AllowMethods = viper.GetString("ALLOWED_METHODS")
	Cors.AllowHeaders = viper.GetString("ALLOWED_HEADERS")
	Cors.ContentLength = viper.GetString("CONTENT_LENGTH")
	Cors.MaxAge = viper.GetInt("MAX_AGE")
	Cors.AllowCredentials = viper.GetBool("ALLOW_CREDENTIALS")

	Mail = &MailConfig{}
	Mail.SMTPHost = viper.GetString("STMP_HOST")
	Mail.SMTPPort = viper.GetInt("STMP_PORT")
	Mail.SenderName = viper.GetString("SENDER_NAME")
	Mail.SenderEmail = viper.GetString("SENDER_EMAIL")
	Mail.SMTPPassword = viper.GetString("STMP_PASSWORD")

	return nil
}
