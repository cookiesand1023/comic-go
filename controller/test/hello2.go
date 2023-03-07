package test

import (
	"encoding/json"
	"net/http"
)

var items []*ItemParams

type ItemParams struct {
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
	Sex  string `json:"sex,omitempty"`
}

func Hello2(w http.ResponseWriter, r *http.Request) {
	//reqBody, _ := io.ReadAll(r.Body)
	var item ItemParams
	//if err := json.Unmarshal(reqBody, &item); err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(item.Age)
	//
	//items = append(items, &item)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(item)
	if err != nil {
		return
	}
}
