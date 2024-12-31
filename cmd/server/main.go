package main

import (
	"net/http"

	"github.com/EXClub/test_project.git/api"
	"github.com/go-chi/chi"
)

func main() {
	//создаём роутер
	r := chi.NewRouter()
	//привязываем роутер к маршрутам
	api.RegisterRoutes(r)
	http.ListenAndServe(":8080", r)

}
