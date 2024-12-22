package api

import "net/http"

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// логика главной
}

func authHandler(w http.ResponseWriter, r *http.Request) {
	// логика страницы авторизации
}

func userPageHandler(w http.ResponseWriter, r *http.Request) {
	// логика страницы пользователя
}

func storyHandler(w http.ResponseWriter, r *http.Request) {
	// логика страницы с прошедшими встречами
}
