package db

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

//type dbRepo struct {
//	Db *sql.DB
//	sync.RWMutex
//}

var DB *sql.DB

func NewSQLite3Repo(dbfile string) error {
	db, err := sql.Open("sqlite3", dbfile)
	if err != nil {
		return errors.New("failed to connect to database")
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	if err := db.Ping(); err != nil {
		return err
	}
	createTables(db)
	//return &dbRepo{
	//	Db: db,
	//}, nil
	DB = db

	return nil

}

func createTables(db *sql.DB) {

	createUsersTables := `
 	CREATE TABLE IF NOT EXISTS users (
    	id INTEGER PRIMARY KEY AUTOINCREMENT,
 	    email TEXT NOT NULL UNIQUE,
 	    password TEXT NOT NULL   
   ) `
	_, err := db.Exec(createUsersTables)
	if err != nil {
		fmt.Println(err)
		panic("Could not create users table.")
	}

	createEventsTable := `
 	CREATE TABLE IF NOT EXISTS events (
    	id INTEGER PRIMARY KEY AUTOINCREMENT,
 	    name TEXT NOT NULL,
 	    description TEXT NOT NULL,
 	    location TEXT NOT NULL,
 	    dateTime DATETIME NOT NULL,
 	    user_id INTEGER,
 	    FOREIGN KEY(user_id) REFERENCES users(id)
   ) `
	_, err = db.Exec(createEventsTable)
	if err != nil {
		panic("Could not create events table.")
	}

	createRegistrationsTable := `
 	CREATE TABLE IF NOT EXISTS registrations (
    	id INTEGER PRIMARY KEY AUTOINCREMENT,
    	event_id INTEGER,
    	user_id INTEGER,
 	    FOREIGN KEY(user_id) REFERENCES users(id)
 	    FOREIGN KEY(event_id) REFERENCES events(id)
   ) `
	_, err = db.Exec(createRegistrationsTable)
	if err != nil {
		panic("Could not create registration table.")
	}

}

//func createTables(db *sql.DB) {
//	//createUsersTables := `
//	//CREATE TABLE IF NOT EXISTS users (
//	// 	id INTEGER PRIMARY KEY AUTOINCREMENT,
//	// 	email TEXT NOT NULL UNIQUE,
//	//    password TEXT NOT NULL,
//	//) `
//	//_, err := db.Exec(createUsersTables)
//	//if err != nil {
//	//	panic("Could not create user table.")
//	//}
//
//	createEventsTable := `
// 	CREATE TABLE IF NOT EXISTS events (
//    	id INTEGER PRIMARY KEY AUTOINCREMENT,
// 	    name TEXT NOT NULL,
// 	    description TEXT NOT NULL,
// 	    location TEXT NOT NULL,
// 	    dateTime DATETIME NOT NULL,
// 	    user_id INTEGER
// 	    FOREIGN KEY(user_id) REFERENCES users(id)
//   ) `
//	_, err := db.Exec(createEventsTable)
//	if err != nil {
//		panic("Could not create events table.")
//	}
//}
