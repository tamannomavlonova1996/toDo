package models

import "time"

type Task struct {
	ID        int       `json:"id" gorm:"column:id"`
	UserID    int       `json:"user_id" gorm:"column:user_id"`
	Name      string    `json:"name" gorm:"column:name"`
	Done      bool      `json:"done" gorm:"column:done"`
	CreatedAt time.Time `json:"created_at,omitempty" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty" gorm:"column:updated"`
	Deadline  time.Time `json:"deadline" gorm:"column:deadline"`

	User *User `json:"user,omitempty" gorm:"foreignKey:UserID"`
}
type User struct {
	ID        int       `json:"id" gorm:"column:id"`
	FullName  string    `json:"full_name" gorm:"column:full_name;unique;not null" `
	Email     string    `json:"email" gorm:"column:email;not null"`
	Password  string    `json:"password" gorm:"column:password;not null"`
	CreatedAt time.Time `json:"created_at,omitempty" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty" gorm:"column:updated_at"`
}

type Tasks []*Task

type UserWithTasks struct {
	User  *UserResponse
	Tasks []*Task
}

type UserResponse struct {
	ID       int    `json:"id"`
	FullName string `json:"full_name"`
}
