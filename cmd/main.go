package main

import (
	"github.com/rayhan889/talkz-v2/app"
	"github.com/rayhan889/talkz-v2/app/integrations/database"
	"github.com/rayhan889/talkz-v2/config"
	"github.com/rayhan889/talkz-v2/pkg/logger"
	redisPkg "github.com/rayhan889/talkz-v2/pkg/redis"
	"github.com/redis/go-redis/v9"
	"gopkg.in/gomail.v2"
	"gorm.io/gorm"
)

var db *gorm.DB
var redisClient *redis.Client
var dialer *gomail.Dialer

func init() {
	var err error
	appConf := config.AppConfig{
		Env: config.Envs.App.Env,
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

	mailConfig := config.MailConfig{
		SMTPHost:     config.Envs.Mail.SMTPHost,
		SMTPPort:     config.Envs.Mail.SMTPPort,
		SenderEmail:  config.Envs.Mail.SenderEmail,
		SMTPPassword: config.Envs.Mail.SMTPPassword,
	}

	dialer = gomail.NewDialer(
		mailConfig.SMTPHost,
		mailConfig.SMTPPort,
		mailConfig.SenderEmail,
		mailConfig.SMTPPassword,
	)

	logger.InitLogger(appConf.Env)

	logger.Log.Info("App Started")

	db, err = database.CreateConnection(
		dbConf.Address,
		int(dbConf.MaxOpenConns),
		int(dbConf.MaxIdleConns),
		dbConf.MaxIdleTime,
	)
	if err != nil {
		logger.Log.Fatal(err)
	}

	logger.Log.Info("Database connection established")

	redisClient, err = redisPkg.InitRedisClient(redisConf.Address, redisConf.Password, int(redisConf.DB))
	if err != nil {
		logger.Log.Fatal(err)
	}
	logger.Log.Info("Redis client initialized")
}

func main() {
	sqlDB, _ := db.DB()
	defer sqlDB.Close()
	defer logger.Log.Sync()

	app := app.InitializeApp(db, redisClient, dialer)
	app.Run()
}
