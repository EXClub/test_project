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
	Hash     []byte
	Token    string
	IsAdmin  bool
}

type Сrypter interface {
	HashPassword(password string)
	ComparePassHash(password string, hash []byte)
}

func (u *User) HashPassword(password string) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		return
	}
	u.Hash = hashedPassword
}

func (u *User) ComparePassHash(password string, hash []byte) error {
	err := bcrypt.CompareHashAndPassword(hash, []byte(password))
	if err != nil {
		return err
	} else {
		return nil
	}
}
