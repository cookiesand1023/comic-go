package auth

import (
	"encoding/json"
	"fmt"
	errors "github.com/comic-go/entity"
	"github.com/comic-go/model"
	"golang.org/x/crypto/bcrypt"
	"io"
	"net/http"
	"time"
)

type SignInRequest struct {
	Email    string
	Password string
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := io.ReadAll(r.Body)
	var signInReq SignInRequest

	// MEMO 値が空の場合、エラーにはならず、構造体の値が空になる
	// TODO リクエストバリデーション
	if err := json.Unmarshal(reqBody, &signInReq); err != nil {
		http.Error(w, fmt.Sprintf("%s", err), http.StatusInternalServerError)
		return
	}

	// Email存在チェック
	user, err := model.GetUserByEmail(signInReq.Email)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode(errors.EmailNotFound)
		if err != nil {
			return
		}
		return
	}

	// password検証
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(signInReq.Password))
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode(errors.PasswordInvalid)
		if err != nil {
			return
		}
		return
	}

	expiration := time.Now()
	expiration = expiration.AddDate(0, 0, 1)
	cookie := http.Cookie{Name: "cg", Value: user.IdToken, Expires: expiration}
	http.SetCookie(w, &cookie)

	_, err = io.WriteString(w, "true")
	if err != nil {
		http.Error(w, fmt.Sprintf("%s", err), http.StatusInternalServerError)
		return
	}

}
