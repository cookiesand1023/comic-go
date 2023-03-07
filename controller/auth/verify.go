package auth

import (
	"encoding/json"
	"fmt"
	"github.com/comic-go/config/jwt"
	"github.com/comic-go/model"
	"net/http"
)

func Verify(w http.ResponseWriter, r *http.Request) {
	// cookieからtokenを取得
	cookie, err := r.Cookie("cg")
	if err != nil {
		fmt.Println(err)
		http.Error(w, fmt.Sprintf("%s", err), http.StatusBadRequest)
		return
	}

	// token検証
	uid, err := jwt.VerifyToken(cookie.Value)
	if err != nil {
		http.Error(w, fmt.Sprintf("%s", err), http.StatusInternalServerError)
		return
	}

	user, err := model.GetUserByUuid(uid)
	if err != nil {
		http.Error(w, fmt.Sprintf("%s", err), http.StatusInternalServerError)
		return
	}

	fmt.Println("verified")

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		return
	}
}
