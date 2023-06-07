package repository

import "ginCli/models"

func (s *Storage) migrate() {
	s.DB.AutoMigrate(
		&models.User{},
		&models.Task{},
	)
}
