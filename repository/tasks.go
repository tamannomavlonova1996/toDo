package repository

import (
	"ginCli/models"
	"gorm.io/gorm"
	"log"
)

func (s *Storage) GetTasks() ([]*models.Task, error) {
	var tasks []*models.Task
	err := s.DB.Model(&models.Task{}).Find(&tasks).Error
	if err != nil {
		log.Println("db, GetTasks, err ", err)
		return nil, err
	}
	return tasks, nil

}

func (s *Storage) GetTaskByID(id int) (*models.Task, error) {
	var task *models.Task
	err := s.DB.Model(&models.Task{}).First(&task, "id=?", id).Error
	if err != nil {
		log.Println("db, GetTaskByID, err ", err)
		return nil, err
	}
	return task, nil
}

func (s *Storage) GetTasksByIDUser(usersID int) ([]*models.Task, error) {
	var userTasks []*models.Task
	err := s.DB.Model(&models.Task{}).
		Where("user_id = ?", usersID).
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, full_name")
		}).
		Find(&userTasks).
		Error
	if err != nil {
		log.Println("db, GetTaskByIDUser, err ", err)
		return nil, err
	}
	return userTasks, nil
}

func (s *Storage) AddTask(newTask *models.Task) (int, error) {
	err := s.DB.Model(&models.Task{}).Create(&newTask).Error
	if err != nil {
		log.Println("db, AddTask, err ", err)
		return 0, err
	}
	return newTask.ID, nil
}

func (s *Storage) UpdateTask(id int, newTask *models.Task) error {
	return s.DB.Model(&models.User{}).Where("id=?", id).Save(&newTask).Error
}

func (s *Storage) DeletedTask(id int) error {
	return s.DB.Model(&models.User{}).Delete(&models.User{}, "id=?", id).Error
}
