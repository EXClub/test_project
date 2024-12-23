package model

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

/*
этот файл только описывает пользователя как объект для взаимодействия с БД
*/
type User struct {
	ID       int
	Username string
	hash     []byte
	token    string
	isAdmin  bool
}

type crypter interface {
	hashPassword(password string)
	checkPassAndHash(password string, hash []byte)
}

func (u *User) hashPassword(password string) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		return
	}
	u.hash = hashedPassword
}

func (u *User) checkPassAndHash(password string, hash []byte) error {
	err := bcrypt.CompareHashAndPassword(u.hash, []byte(password))
	if err != nil {
		return err
	} else {
		return nil
	}
}
