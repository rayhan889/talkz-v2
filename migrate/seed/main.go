package main

import (
	"log"

	"github.com/rayhan889/talkz-v2/app/integrations/database"
	"github.com/rayhan889/talkz-v2/config"
)

func main() {
	dbConf := config.DBConfig{
		Address:      config.DB.Address,
		MaxOpenConns: config.DB.MaxOpenConns,
		MaxIdleConns: config.DB.MaxIdleConns,
		MaxIdleTime:  config.DB.MaxIdleTime,
	}
	db, err := database.CreateConnection(
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
