package model

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDatabase() (err error) {
	connectInfo := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
	)

	for i := 0; i < 10; i++ {
		if Db, err = gorm.Open(mysql.Open(connectInfo), &gorm.Config{}); err != nil {
			time.Sleep(time.Second * 2)
			log.Printf("retry...")
			continue
		}
		fmt.Println("Database connected!")
		break
	}

	migrateDatabase()
	// エラーを返す
	return err
}

func migrateDatabase() {
	Db.AutoMigrate(User{})
}
