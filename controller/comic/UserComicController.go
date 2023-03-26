package comic

import (
	"encoding/json"
	"fmt"
	errors "github.com/comic-go/entity"
	"github.com/comic-go/model"
	"io"
	"net/http"
	"strconv"
)

func GetComicsIsRead(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query().Get("uid")
	userId, err := strconv.Atoi(param)
	if err != nil {
		fmt.Println(err)
		http.Error(w, fmt.Sprintf("%s", err), http.StatusBadRequest)
		return
	}
	comics, err := model.UserComicsIsReadByUserId(userId)
	if err != nil {
		fmt.Println(err)
		http.Error(w, fmt.Sprintf("%s", err), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(comics)
	if err != nil {
		return
	}
}

type RegisterUserComicStatusRequest struct {
	UserId  int    `json:"user_id"`
	ComicId int    `json:"comic_id"`
	Type    string `json:"type"`
	Status  bool   `json:"status"`
}

func RegisterUserComicStatus(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := io.ReadAll(r.Body)
	var req RegisterUserComicStatusRequest

	// MEMO 値が空の場合、エラーにはならず、構造体の値が空になる
	// TODO リクエストバリデーション
	if err := json.Unmarshal(reqBody, &req); err != nil {
		http.Error(w, fmt.Sprintf("%s", err), http.StatusInternalServerError)
		return
	}

	if req.Type != "is_read" && req.Type != "will_read" {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusInternalServerError)
		err := json.NewEncoder(w).Encode(errors.TypeInvalid)
		if err != nil {
			fmt.Println(err)
			return
		}
		return
	}

	_, err := model.UpsertUserComic(req.UserId, req.ComicId, req.Type, req.Status)
	if err != nil {
		fmt.Println(err)
		http.Error(w, fmt.Sprintf("%s", err), http.StatusBadRequest)
		return
	}

	_, err = io.WriteString(w, "true")
	if err != nil {
		http.Error(w, fmt.Sprintf("%s", err), http.StatusBadRequest)
		return
	}
}
