package auth

import (
	"fmt"
	"net/http"

	"github.com/comic-go/config/jwt"
)

func Test(w http.ResponseWriter, _ *http.Request) {

	token, err := jwt.CreateToken("subject")
	if err != nil {
		http.Error(w, fmt.Sprintf("%s", err), http.StatusInternalServerError)
	}

	userName, err := jwt.VerifyToken(token)
	if err != nil {
		http.Error(w, fmt.Sprintf("%s", err), http.StatusInternalServerError)
	}

	_, err = fmt.Fprintln(w, "JWT Auth Complete!!"+userName)
	if err != nil {
		http.Error(w, fmt.Sprintf("%s", err), http.StatusInternalServerError)
	}
}
