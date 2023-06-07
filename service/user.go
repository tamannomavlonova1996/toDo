package service

import (
	"ginCli/models"
)

func (s *Service) CreateUser(u models.User) (int, error) {
	u.Password = generatePasswordHash(u.Password)
	return s.Storage.CreateUser(&u)
}
