package model

type User struct {
	id       int
	username string
	hash     []byte
	token    string
	isAdmin  bool
}
