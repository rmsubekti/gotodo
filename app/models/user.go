package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
}

func (u *User) Register() error {

	return nil
}

func (u *User) Login() error {

	return nil
}
