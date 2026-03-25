package models

import (
	"errors"
	"learning/rest-api/db"
	"learning/rest-api/utils"
)

type User struct {
	Id       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	query := `insert into users(email,password) values(?,?)`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()
	hasPass, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(u.Email, hasPass)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()

	u.Id = id
	return err
}

func (u *User) ValidateUser() error {
	query := `select id, password from users where email=?`

	row := db.DB.QueryRow(query, u.Email)

	var dbPass string

	err := row.Scan(&u.Id, &dbPass)
	if err != nil {
		return errors.New("db err")
	}

	passmatch := utils.MathchPassword(u.Password, dbPass)

	if !passmatch {
		return errors.New("invalid credentials")
	}

	return nil
}
