package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Sex      string `json:"sex"`
	Email    string `gorm:"type:varchar(100);unique_index" json:"email"`
	Uuid     string `gorm:"type:varchar(100);unique_index" json:"uuid"`
	IdToken  string `json:"id_token"`
	Password string `json:"password"`
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
