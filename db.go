package main

import (
  "database/sql"
  _ "github.com/mattn/go-sqlite3"
  "log"
  "sync"
)

var (
  DB        *sql.DB
  dbInitErr error
  dbMutex   sync.Mutex
)

const file string = "./db.db" // Path to the SQLite database file

func ConnectDb() error {
  // Connects to the SQLite database and returns an error if it fails.

  dbMutex.Lock() // lock on call
  defer dbMutex.Unlock()

  db, err := sql.Open("sqlite3", file)
  if err != nil {
    log.Fatal(err)
    return err
  }
  DB = db

  return nil
}

func DisconnectDb() {
	dbMutex.Lock() // lock on call
	defer dbMutex.Unlock()
  
	if DB != nil {
	  DB.Close()
	}
}