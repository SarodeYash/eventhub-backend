package models

import (
	"errors"

	util "main.go/Util"
	"main.go/db"
)

type Users struct {
	UserId   int64
	Email    string `binding:"required"` //used to specify this are the field that are must required
	Password string `binding:"required"`
}

// The function is used for the New user registration
func (u *Users) RegisterNewUser() error {
	query := `
	INSERT INTO "user"(Email,Password)
	VALUES($1,$2)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	Password, err := util.ToHahshPass(u.Password) //This will hash our plane password and stores that password in Database
	if err != nil {
		return err
	}
	_, err = stmt.Exec(u.Email, Password)

	return err
	// Last_Id, err := result.LastInsertId() //Each user will have an ID
	// u.UserId = Last_Id
	// return err

}

// This function validates password for the given mail with the user data stored in database
func (u *Users) Login() error {
	query := `
	SELECT id,Password FROM "user" WHERE Email=$1
	`
	row := db.DB.QueryRow(query, u.Email)
	var retrivedPassword string
	err := row.Scan(&u.UserId, &retrivedPassword) //Used to scan/get a single value of password from the database
	if err != nil {
		return errors.New("invalid password")
	}
	ValidPassword := util.CompareHashPassword(u.Password, retrivedPassword) //this method is used to compare the user provided password and thr password stored in the database
	if !ValidPassword {
		return errors.New("invalid password")
	}
	return nil

}
