package api

import (
	"net/http"

	"github.com/go-chi/chi"
)

func registerRoutes(r *chi.Mux) {

	//главная
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("home"))
	})

	//авторизация
	r.Get("/auth", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("home"))
	})

	//личный кабинет
	r.Get("/user", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("home"))
	})

	//история встреч
	r.Get("/story", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("home"))
	})

}
