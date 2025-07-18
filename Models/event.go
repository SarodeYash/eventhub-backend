package models

import (
	"time"

	"main.go/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int64
}

// Will used to save the data for an event that is been shared by the client with the Post Method
func (e *Event) Save() error {
	query := `INSERT INTO event (name,description,location,DateTime,User_Id)
	Values($1,$2,$3,$4,$5);` //Data will be inserted in SQL database  with insert query

	stmt, err := db.DB.Prepare(query) //Preapare method is similar to Exec but mostly used when we want to prepare a query and then after use it with some args
	if err != nil {
		return err
	}

	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID) //Exec is used when we need to insert delete or update database
	// id, err := result.LastInsertId() //This will provide as LastInserted ID which will auto incerement as mentioned in db.go
	// e.ID = id                        //Which we will store in current ID
	defer stmt.Close() //after copleting this save() method staement need to be closed
	// return err
	return err
}

// All the event data will be displayed
func GetAllEvents() ([]Event, error) {
	query := `SELECT *FROM event`   //Query to select or get all the database
	rows, err := db.DB.Query(query) //Query is used to fetch the data from the databse which is stored in rows
	if err != nil {
		return nil, err
	}
	defer rows.Close()      //rows needs to be closed after completion of GetAllEvents()
	var Empty_Event []Event //Creating Empty slice
	for rows.Next() {       //This method will iterate for loop till there's no mext row
		var Old_Event Event
		err := rows.Scan(&Old_Event.ID, &Old_Event.Name, &Old_Event.Description, &Old_Event.Location, &Old_Event.DateTime, &Old_Event.UserID) //Fetching data into the Slice using SCan method which copies column data to the mentioned var
		if err != nil {
			return nil, err
		}

		Empty_Event = append(Empty_Event, Old_Event) //Adding the fetched data to the empty slice
	}
	return Empty_Event, nil
}

func GetIDEvent(Id int64) (*Event, error) {
	query := `SELECT *FROM event WHERE Id = $1`
	rows_id := db.DB.QueryRow(query, Id) //QueryRow Will get a single row data

	var event Event
	err := rows_id.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (event Event) DeleteEvent() error {
	query := `DELETE FROM event WHERE Id = $1` //Delete Query to Delete Event mentioned ID
	_, err := db.DB.Exec(query, event.ID)
	return err
}

func (event Event) UpdateEvent() error {
	query := `
	UPDATE event
	SET Name=$1,Description=$2,Location=$3,DateTime=$4
	WHERE Id=$5`
	upState, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer upState.Close()
	_, err = upState.Exec(event.Name, event.Description, event.Location, event.DateTime, event.ID)
	return err

}
