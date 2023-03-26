package model

import (
	"gorm.io/gorm"
)

type Comic struct {
	gorm.Model
	Title    string `gorm:"unique_index" json:"title"`
	Author   string `json:"author"`
	Volume   int    `json:"volume"`
	ImageUrl string `json:"image_url"`
}

var comic Comic
var comics []Comic

func GetAllComics() (res []Comic, err error) {
	err = Db.Find(&comics).Error
	return comics, err
}

func GetComicById(id int) (res Comic, err error) {
	err = Db.Find(&comic, id).Error
	return comic, err
}
