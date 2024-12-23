package model

/*
этот файл только описывает книгу как объект для взаимодействия с БД
*/
type Book struct {
	ID     int
	Name   string
	Author string
	Genre  string
}
