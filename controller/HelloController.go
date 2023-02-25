package controller

import (
	"fmt"
	"net/http"

	"github.com/comic-go/model"
)

func HelloController(w http.ResponseWriter, r *http.Request) {
	user, err := model.GetFirst()
	if err != nil {
		http.Error(w, fmt.Sprintf("%s", err), http.StatusInternalServerError)
	}

	_, err = fmt.Fprintln(w, "Hello, World!"+user.Name)
	if err != nil {
		http.Error(w, fmt.Sprintf("%s", err), http.StatusInternalServerError)
	}
}
