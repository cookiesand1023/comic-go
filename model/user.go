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
	Uuid     string `gorm:"type:varchar(100);unique_index"`
	IdToken  string
	Password string
}

func GetFirst() (user User, err error) {
	user = User{}
	err = Db.First(&user).Error
	return user, err
}

func GetUserByEmail(email string) (user User, err error) {
	user = User{}
	err = Db.First(&user, "email = ?", email).Error
	return user, err
}

func GetUserByUuid(uuid string) (user User, err error) {
	user = User{}
	err = Db.First(&user, "uuid = ?", uuid).Error
	return user, err
}

func UpdateToken(uuid string, token string) (user User, err error) {
	user = User{}
	err = Db.First(&user, "uuid = ?", uuid).Update("id_token", token).Error
	return user, err
}
