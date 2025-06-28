package config

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/rayhan889/talkz-v2/app/env"
)

type Config struct {
	App   AppConfig
	DB    DBConfig
	Redis RedisConfig
	JWT   JWTConfig
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

var Envs = initEnvs()

func initEnvs() *Config {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error loading .env file: %v\n", err)
	}

	return &Config{
		App: AppConfig{
			Port:    env.GetString("PORT", ":8080"),
			Env:     env.GetString("ENV", "development"),
			Version: env.GetString("VERSION", "v1.0.0"),
		},
		DB: DBConfig{
			Address:      env.GetString("DB_ADDR", "postgres://postgres:112233@localhost:5431/talkzdb?sslmode=disable"),
			MaxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 10),
			MaxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 5),
			MaxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "5m"),
		},
		Redis: RedisConfig{
			Address:  env.GetString("REDIS_ADDR", "localhost:6379"),
			Password: env.GetString("REDIS_PASSWORD", ""),
			DB:       env.GetInt("REDIS_DB", 0),
			Protocol: env.GetInt("REDIS_PROTOCOL", 2),
		},
		JWT: JWTConfig{
			Secret:  env.GetString("JWT_SECRET", "secret"),
			Expires: env.GetInt("JWT_EXPIRATIONS_IN_SECOND", 3600),
		},
	}
}
