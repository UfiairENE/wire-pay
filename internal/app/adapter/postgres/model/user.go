package model

import (
	"time"
)

//Tabler is the interface of GORM table name
type Tabler interface {
	TableName() string
}

type User struct {
	// gorm.Model
	ID          uint
	FirstName   string    `gorm:"type:text; not null"`
	LastName    string    `gorm:"type:text; not null"`
	Email       string    `gorm:"type:text; not null"`
	PhoneNumber string    `gorm:"type:text; not null"`
	CreatedAt   time.Time `gorm:"created_at, not null"`
	UpdatedAt   time.Time `gorm:"updated_at, not null"`
}

//TableName gets table name of user
func (User) TableName() string {
	return "user"
}
