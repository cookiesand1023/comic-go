package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	// 環境未設定の場合はローカルを設定
	if "" == os.Getenv("GO_ENV") {
		_ = os.Setenv("GO_ENV", "development")
	}
	err := godotenv.Load(fmt.Sprintf(".env.%s", os.Getenv("GO_ENV")))
	if err != nil {
		log.Fatalf("Error loading env target env is %s", os.Getenv("GO_ENV"))
	}
}
