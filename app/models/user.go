package models

import (
	"errors"

	"github.com/segmentio/ksuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       string `gorm:"type:varchar(30);primarykey:true;not null;unique" json:"id"`
	Name     string `gorm:"type:varchar(30);not null" json:"name"`
	Email    string `gorm:"type:varchar(30);not null;unique" json:"email"`
	Password string `gorm:"type:varchar(200);not null" json:"password"`
}

func (u *User) Register() error {
	u.ID = ksuid.New().String()

	// save hashed password to db
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	u.Password = string(hashedPassword)
	return db.Create(&u).Error
}

func (u *User) Login(email, password string) error {
	result := db.First(&u, "email = ?", email)
	if result.RowsAffected < 1 {
		return errors.New("email not registered")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return errors.New("wrong password")
	}
	return nil
}
