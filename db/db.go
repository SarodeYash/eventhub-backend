package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB //Setting handler for the SQL with sql pointer

// Setting up the database Connection
func InitDB() {
	var err error
	DB, err = sql.Open("postgres", "user=postgres password=yash7807 dbname=Event_API host=localhost port=5432 sslmode=disable") //Setting Driver as an postgresql and metining path for the database with open method
	if err != nil {
		log.Fatal("Failed to open Database") //this method will exit the code if err is there
	}
	err = DB.Ping()
	if err != nil {
		log.Fatal("Error in connecting Database")
	}

	DB.SetMaxOpenConns(10) //This method will keep 10 connections open for the database
	DB.SetMaxIdleConns(5)  //of which 5 are continuesly alive and rest 5 will be if required
	CreateTable()
}

// Creating table for the Data in Database
func CreateTable() {
	SignUpTable := `
	CREATE TABLE IF NOT EXISTS "user"(
	
	Id SERIAL PRIMARY KEY ,
	Email TEXT NOT NULL UNIQUE,
	Password TEXT NOT NULL	
	)
	
	`
	_, err := DB.Exec(SignUpTable)
	if err != nil {
		log.Fatal("User Table not Created")
	}
	CreateEventTable := `
	
	CREATE TABLE IF NOT EXISTS event(
	
	Id SERIAL PRIMARY KEY,
	Name TEXT NOT NULL,
	Description TEXT NOT NULL,
	Location TEXT NOT NULL,
	DateTime TIMESTAMP NOT NULL,
	User_Id INTEGER,
	FOREIGN KEY (User_Id) REFERENCES "user"(Id)

	);`
	//Query for creating the Database
	_, err = DB.Exec(CreateEventTable) //Executes an mentioned sql query
	if err != nil {
		log.Fatal("Event Table not Created")
	}

	Register := `
	CREATE TABLE IF NOT EXISTS Register(

	Id SERIAL PRIMARY KEY ,
	event_id INTEGER , 
	user_id INTEGER , 
	FOREIGN KEY (Event_id) REFERENCES event(Id),
	FOREIGN KEY (User_id) REFERENCES "user"(Id),
	UNIQUE(event_id,user_id)

	);`
	_, err = DB.Exec(Register)
	if err != nil {
		log.Fatal("Registarion Table Not Created")

	}

}
