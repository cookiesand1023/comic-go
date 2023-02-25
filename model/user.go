package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Age      int
	Sex      string
	Email    string `gorm:"type:varchar(100);unique_index"`
	Password string
}

func GetFirst() (user User, err error) {
	result := User{}
	err = db.First(&result).Error
	return result, err
}
