package main

import (
	"log"

	"github.com/rayhan889/talkz-v2/config"
	database "github.com/rayhan889/talkz-v2/internal/db"
)

func main() {
	dbConf := config.DBConfig{
		Address:      config.Envs.DB.Address,
		MaxOpenConns: config.Envs.DB.MaxOpenConns,
		MaxIdleConns: config.Envs.DB.MaxIdleConns,
		MaxIdleTime:  config.Envs.DB.MaxIdleTime,
	}
	db, err := database.New(
		dbConf.Address,
		int(dbConf.MaxOpenConns),
		int(dbConf.MaxIdleConns),
		dbConf.MaxIdleTime,
	)
	if err != nil {
		log.Fatal(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()

	database.Seeder(db)
}
