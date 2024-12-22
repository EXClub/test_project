package api

import (
	"github.com/go-chi/chi"
)

// функция, которая создаёт маршруты(пути)
func RegisterRoutes(r *chi.Mux) {

	//главная
	r.Get("/", HomeHandler)

	//авторизация
	r.Get("/auth", AuthHandler)

	//личный кабинет
	r.Get("/user", UserPageHandler)

	//история встреч
	r.Get("/story", StoryHandler)

}
