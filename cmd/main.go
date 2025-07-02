package main

import (
	"log"

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

	err = config.LoadConfig()

	if err != nil {
		log.Fatal(err)
	}

	dialer = gomail.NewDialer(
		config.Mail.SMTPHost,
		config.Mail.SMTPPort,
		config.Mail.SenderEmail,
		config.Mail.SMTPPassword,
	)

	logger.InitLogger(config.App.Env)

	logger.Log.Info("App Started")

	db, err = database.CreateConnection(
		config.DB.Address,
		int(config.DB.MaxOpenConns),
		int(config.DB.MaxIdleConns),
		config.DB.MaxIdleTime,
	)
	if err != nil {
		logger.Log.Fatal(err)
	}

	logger.Log.Info("Database connection established")

	redisClient, err = redisPkg.InitRedisClient(config.Redis.Address, config.Redis.Password, int(config.Redis.DB))
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
