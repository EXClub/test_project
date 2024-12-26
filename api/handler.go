package api

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// логика главной
	fmt.Fprintf(w, "Hello World!")
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
