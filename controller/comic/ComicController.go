package comic

import (
	"encoding/json"
	"fmt"
	"github.com/comic-go/model"
	"net/http"
	"strconv"
)

func GetAllComics(w http.ResponseWriter, r *http.Request) {
	comics, err := model.GetAllComics()

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(comics)
	if err != nil {
		return
	}
}

func GetComicDetail(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query().Get("cid")
	cid, err := strconv.Atoi(param)
	if err != nil {
		fmt.Println(err)
		http.Error(w, fmt.Sprintf("%s", err), http.StatusBadRequest)
		return
	}
	comic, err := model.GetComicById(cid)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(comic)
	if err != nil {
		return
	}
}
