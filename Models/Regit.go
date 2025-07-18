package models

import (
	"main.go/db"
)

type Regit struct {
	Id       int64
	Event_Id int64 `binding:"required"`
	User_id  int64 `binding:"required"`
}

func (e *Event) Registartion(userid int64) error {

	query := `INSERT INTO Register (event_id,user_id)
	Values($1,$2)
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(e.ID, userid)
	stmt.Close()
	return err
}
func (e *Event) Cancelation(userId int64) error {
	query := `DELETE FROM Register WHERE event_id=$1 AND user_id=$2`
	_, err := db.DB.Exec(query, e.ID, userId)
	return err
}
