package models

import (
	"errors"
	"learning/rest-api_gorm/db"
	"learning/rest-api_gorm/utils"
)

type User struct {
	Id       int64   `gorm:"primaryKey"`
	Email    string  `binding:"required"  gorm:"unique ; not null"`
	Password string  `binding:"required" gorm:"not null"`
	Events   []Event `gorm:"foreignKey:UserID"`
}

func (u User) Save() error {
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = hashedPassword

	return db.DB.Create(&u).Error

}

func (u *User) ValidateUser() error {
	var dbUser User
	err := db.DB.Where("email = ?", u.Email).First(&dbUser).Error
	if err != nil {
		return errors.New("invalid credentials")
	}

	passmatch := utils.MathchPassword(u.Password, dbUser.Password)

	if !passmatch {
		return errors.New("invalid credentials")
	}
	u.Id = dbUser.Id

	return nil
}
