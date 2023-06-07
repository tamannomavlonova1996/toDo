package repository

import (
	"ginCli/models"
	"log"
)

func (s *Storage) GetUsers() ([]*models.User, error) {
	var users []*models.User

	err := s.DB.Model(&models.User{}).Find(&users).Error
	if err != nil {
		log.Println("db, GetUsers, err ", err)
		return nil, err
	}
	return users, nil
}

func (s *Storage) CreateUser(newUser *models.User) (int, error) {
	err := s.DB.Model(&models.User{}).Create(&newUser).Error
	if err != nil {
		log.Println("db, CreateUser, err ", err)
		return 0, err
	}
	return newUser.ID, nil
}

func (s *Storage) UpdateUser(id int, newUser *models.User) error {
	return s.DB.Model(&models.User{}).Where("id=?", id).Save(&newUser).Error
}

func (s *Storage) DeletedUser(id int) error {
	return s.DB.Model(&models.User{}).Delete(&models.User{}, "id=?", id).Error
}

func (s *Storage) GetUserByID(id int) (*models.User, error) {
	var user *models.User
	err := s.DB.Model(&models.User{}).First(&user, "id=?", id).Error
	if err != nil {
		log.Println("db, GetUserByID, err ", err)
		return nil, err
	}
	return user, nil
}

func (s *Storage) GetUserByUserNameAndPass(username, password string) (*models.User, error) {
	var user *models.User
	err := s.DB.Model(&models.User{}).First(&user, "full_name=? AND password=?", username, password).Error
	if err != nil {
		log.Println("db, GetUsers, err ", err)
		return nil, err
	}
	return user, nil
}
