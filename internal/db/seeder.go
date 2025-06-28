package db

import (
	"log"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/rayhan889/talkz-v2/internal/model"
	"github.com/rayhan889/talkz-v2/pkg/uuid"
	"gorm.io/gorm"
)

type Seed struct {
	db *gorm.DB
}

func Seeder(db *gorm.DB) error {
	seed := &Seed{db: db}
	log.Println("Seeding database...")

	if err := seed.seedUsers(); err != nil {
		log.Fatalf("Error seeding users: %v", err)
		return err
	}

	log.Println("Database seeding completed successfully.")
	return nil
}

func (s *Seed) seedUsers() error {
	var users []model.User
	for i := 0; i < 10; i++ {
		user := model.User{
			ID:        uuid.GenerateUUID(),
			Username:  faker.Username(),
			Email:     faker.Email(),
			Password:  faker.Password(),
			CreatedAt: time.Now(),
		}
		users = append(users, user)
	}
	err := s.db.Create(&users).Error
	if err != nil {
		return err
	}
	return nil
}
