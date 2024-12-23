package model

/*
этот файл только описывает пользователя как объект для взаимодействия с БД
*/
type user struct {
	ID       int
	Username string
	Password string
}
