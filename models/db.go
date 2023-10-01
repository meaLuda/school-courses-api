package models

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	// "sync"
)

var (
	DB        *sql.DB
	dbInitErr error
	// dbMutex   sync.Mutex
)
const file string = "./db.sqlite3"


func ConnectDb() error {
	// dbMutex.Lock() // lock on call
	// defer dbMutex.Unlock()

	db, err := sql.Open("sqlite3", file)
	
	if err != nil {
		log.Fatal(err)
		return err
	}
	DB = db
	return nil
	
}