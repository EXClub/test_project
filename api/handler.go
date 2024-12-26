package api

import (
	"fmt"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// логика главной
	fmt.Fprintf(w, "Hello World!")
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
