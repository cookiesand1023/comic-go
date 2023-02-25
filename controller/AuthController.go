package controller

import (
	"fmt"
	"net/http"

	"github.com/comic-go/config/jwt"
)

func AuthController(w http.ResponseWriter, r *http.Request) {

	token, err := jwt.CreateToken()
	if err != nil {
		http.Error(w, fmt.Sprintf("%s", err), http.StatusInternalServerError)
	}

	user_name, err := jwt.VerifyToken(token)
	if err != nil {
		http.Error(w, fmt.Sprintf("%s", err), http.StatusInternalServerError)
	}

	_, err = fmt.Fprintln(w, "JWT Auth Complete!!"+user_name)
	if err != nil {
		http.Error(w, fmt.Sprintf("%s", err), http.StatusInternalServerError)
	}
}
