package config

import (
	"net/http"

	"github.com/comic-go/controller"
)

// setRouter ルーティングをセット
func SetRouter() {
	http.HandleFunc("/test", controller.HelloController)

	http.HandleFunc("/auth", controller.AuthController)
}
