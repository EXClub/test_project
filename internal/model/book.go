package model

/*
этот файл только описывает книгу как объект для взаимодействия с БД
*/
type book struct {
	ID     int
	Name   string
	Author string
	Genre  string
}
