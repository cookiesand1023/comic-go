package config

import (
	"github.com/comic-go/controller/comic"
	"net/http"

	"github.com/comic-go/controller/auth"
	"github.com/comic-go/controller/test"
	"github.com/gorilla/mux"
)

// SetRouter setRouter ルーティングをセット
func SetRouter() {

	r := mux.NewRouter()
	r.Use(CORSMiddleware)

	// test router
	r.HandleFunc("/test", test.Hello)
	testRouter := r.PathPrefix("/test").Subrouter()
	testRouter.HandleFunc("/post", test.Hello2)

	// auth router
	r.HandleFunc("/auth", auth.Test)
	authRouter := r.PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("/verify", auth.Verify).Methods("GET")

	r.HandleFunc("/signup", auth.SignUp)
	r.HandleFunc("/signin", auth.SignIn)

	r.HandleFunc("/comics", comic.GetAllComics).Methods("GET")
	r.HandleFunc("/comic", comic.GetComicDetail).Methods("GET")

	r.HandleFunc("/user_comic/is_read", comic.GetComicsIsRead).Methods("GET")
	r.HandleFunc("/user_comic/register", comic.RegisterUserComicStatus)

	http.Handle("/", r)
}

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := "http://localhost:3001"
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Request-Method", "GET,PUT,POST,DELETE,UPDATE,OPTIONS")
		w.Header().Set("Content-Type", "application/json")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}
