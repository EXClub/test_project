package api

import (
	"github.com/go-chi/chi"
)

// функция, которая создаёт маршруты(пути)
func RegisterRoutes(r *chi.Mux) {

	//главная
	r.Get("/", homeHandler)

	//авторизация
	r.Get("/auth", authHandler)

	//личный кабинет
	r.Get("/user", userPageHandler)

	//история встреч
	r.Get("/story", storyHandler)

}
