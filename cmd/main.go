package main

import (
	"github.com/rayhan889/talkz-v2/config"
	"github.com/rayhan889/talkz-v2/internal/db"
	"github.com/rayhan889/talkz-v2/internal/server"
	"github.com/rayhan889/talkz-v2/pkg/logger"
	"github.com/rayhan889/talkz-v2/pkg/redis"
)

func main() {
	appConf := config.AppConfig{
		Port: config.Envs.App.Port,
		Env:  config.Envs.App.Env,
	}

	dbConf := config.DBConfig{
		Address:      config.Envs.DB.Address,
		MaxOpenConns: config.Envs.DB.MaxOpenConns,
		MaxIdleConns: config.Envs.DB.MaxIdleConns,
		MaxIdleTime:  config.Envs.DB.MaxIdleTime,
	}

	redisConf := config.RedisConfig{
		Address:  config.Envs.Redis.Address,
		Password: config.Envs.Redis.Password,
		DB:       config.Envs.Redis.DB,
	}

	logger.InitLogger(appConf.Env)
	defer logger.Log.Sync()

	logger.Log.Info("App Started")

	db, err := db.New(
		dbConf.Address,
		int(dbConf.MaxOpenConns),
		int(dbConf.MaxIdleConns),
		dbConf.MaxIdleTime,
	)
	if err != nil {
		logger.Log.Fatal(err)
	}

	sqlDB, _ := db.DB()

	defer sqlDB.Close()
	logger.Log.Info("Database connection established")

	err = redis.InitRedisClient(redisConf.Address, redisConf.Password, int(redisConf.DB))
	if err != nil {
		logger.Log.Fatal(err)
	}
	logger.Log.Info("Redis client initialized")

	server := server.NewAPIServer(appConf.Port, db)
	err = server.Start(appConf.Env)
	if err != nil {
		panic(err)
	}

}
