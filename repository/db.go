package repository

import (
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Storage struct {
	DB *gorm.DB
}

func NewStorage() (*Storage, error) {
	dsn := "host=localhost user=user password=pass dbname=todo " +
		"port=5460 sslmode=disable"
	dbGorm, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	db := &Storage{dbGorm}

	db.migrate()
	
	return db, nil
}
