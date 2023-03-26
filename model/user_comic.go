package model

import (
	"errors"
	"gorm.io/gorm"
)

type UserComic struct {
	gorm.Model
	UserId   int  `json:"user_id"`
	ComicId  int  `json:"comic_id"`
	IsRead   bool `json:"is_read" gorm:"default:false"`
	WillRead bool `json:"will_read" gorm:"default:false"`
}

type UserComicDetails struct {
	UserComic
	Comic
}

//var userComics []UserComic

var userComicDetails []UserComicDetails

func UserComicsIsReadByUserId(id int) (res []UserComicDetails, err error) {
	err = Db.Model(&UserComic{}).
		Select("user_comics.*, comics.*").
		Joins("INNER JOIN comics ON user_comics.comic_id = comics.id").
		Where("user_comics.is_read = ?", true).
		Where("user_comics.user_id = ?", id).
		Find(&userComicDetails).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		var emptyArr []UserComicDetails
		return emptyArr, err
	} else {
		return userComicDetails, err
	}
}

func UpsertUserComic(uid int, cid int, t string, status bool) (result bool, err error) {
	//err = Db.Where("user_id = ?", uid).Where("comic_id = ?", cid).Error
	var userComic UserComic
	err = Db.FirstOrCreate(&userComic, UserComic{UserId: uid, ComicId: cid}).Error

	if t == "is_read" {
		err = Db.Model(&userComic).Update("is_read", status).Error
	} else if t == "will_read" {
		err = Db.Model(&userComic).Update("will_read", status).Error
	}
	return true, err
}
