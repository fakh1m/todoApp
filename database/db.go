package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

//membuat koneksi database

var DB *sql.DB

func ConnectDB() {
	var err error
	connStr := "user=kimbozy password=password dbname=todoApp sslmode=disable "
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database Connected Successfully")

}
