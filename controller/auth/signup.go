package auth

import (
	"encoding/json"
	"fmt"
	errors "github.com/comic-go/entity"
	"gorm.io/gorm"
	"io"
	"net/http"
	"time"

	"github.com/comic-go/config/jwt"
	"github.com/comic-go/model"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type SignUpRequest struct {
	Email    string
	UserName string `json:"user_name"`
	Password string
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := io.ReadAll(r.Body)
	var signUpReq SignUpRequest

	// MEMO 値が空の場合、エラーにはならず、構造体の値が空になる
	// TODO リクエストバリデーション
	if err := json.Unmarshal(reqBody, &signUpReq); err != nil {
		http.Error(w, fmt.Sprintf("%s", err), http.StatusInternalServerError)
		return
	}

	// email 重複チェック
	_, err := model.GetUserByEmail(signUpReq.Email)
	if err != gorm.ErrRecordNotFound {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(errors.DuplicateEmailError)
		err := json.NewEncoder(w).Encode(errors.DuplicateEmailError)
		if err != nil {
			fmt.Println(err)
			return
		}
		return
	}

	// uuid生成
	uid, err := createUuid()
	if err != nil {
		http.Error(w, fmt.Sprintf("%s", err), http.StatusInternalServerError)
		return
	}

	// password hash
	hashed, err := hashPassword(signUpReq.Password)
	if err != nil {
		http.Error(w, fmt.Sprintf("%s", err), http.StatusInternalServerError)
		return
	}

	// idToken生成
	token, err := jwt.CreateToken(uid)
	if err != nil {
		http.Error(w, fmt.Sprintf("%s", err), http.StatusInternalServerError)
		return
	}

	user := &model.User{
		Name:     signUpReq.UserName,
		Email:    signUpReq.Email,
		Password: hashed,
		Uuid:     uid,
		IdToken:  token,
	}

	// レコード追加
	result := model.Db.Create(&user)
	if result.Error != nil {
		http.Error(w, fmt.Sprintf("%s", result.Error), http.StatusInternalServerError)
		return
	}

	expiration := time.Now()
	expiration = expiration.AddDate(0, 0, 1)
	cookie := http.Cookie{Name: "cg", Value: token, Expires: expiration}
	http.SetCookie(w, &cookie)

	_, err = io.WriteString(w, uid)
	if err != nil {
		http.Error(w, fmt.Sprintf("%s", err), http.StatusInternalServerError)
		return
	}
}

func createUuid() (string, error) {
	u, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	uu := u.String()
	return uu, err
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
