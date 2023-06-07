package service

import "ginCli/repository"

type Service struct {
	*repository.Storage
}

func NewService(storage *repository.Storage) *Service {
	return &Service{
		Storage: storage,
	}
}
