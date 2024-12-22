package main

import (
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	//создаём роутер
	r := chi.NewRouter()
	//привязываем роутер к маршрутам
	routes.RegisterRoutes(r)

	http.ListenAndServe(":8080", r)
}
