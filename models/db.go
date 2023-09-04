package models

import (
	"database/sql"
	"log"
	"sync"
)

// // global variable
// var DB *sql.DB

// func ConnectDb() error{
// 	db, err := sql.Open("sqlite3","models/db.sqlite3")

// 	// check when an error occurs
// 	if err != nil{
// 		log.Fatal(err)
// 		return err
// 	}

// 	// Set value of DB to this new db
// 	DB = db 
// 	log.Println("--------- Connected to the database ------------------")
// 	return nil
// }

var (
	DB        *sql.DB
	dbInitErr error
	dbMutex   sync.Mutex
)

func ConnectDb() error {
	dbMutex.Lock() // lock on call
	defer dbMutex.Unlock()

	if DB != nil {
		return nil // Already connected
	}

	db, err := sql.Open("sqlite3", "models/db.sqlite3")
	if err != nil {
		log.Fatal(err)
		return err
	}

	DB = db
	log.Println("--------- Connected to the database ------------------")
	return nil
}