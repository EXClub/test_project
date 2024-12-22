package api

import "net/http"

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// логика главной
}

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	// логика страницы авторизации
}

func UserPageHandler(w http.ResponseWriter, r *http.Request) {
	// логика страницы пользователя
}

func StoryHandler(w http.ResponseWriter, r *http.Request) {
	// логика страницы с прошедшими встречами
}
