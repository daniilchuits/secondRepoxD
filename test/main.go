package main

import (
	"database/sql"
	"log"
	"test/internal/handlers"

	_ "github.com/lib/pq"
)

func main() {

	cnnStr := "user=secret-user password=secret-password dbname=new-db host=postgres-db sslmode=disable"
	db, err := sql.Open("postgres", cnnStr)
	if err != nil {
		log.Fatal("Error in opening db:", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("Error pinging db:", err)
	}

	dbManager := handlers.NewDB(db)
	if err = dbManager.Create(); err != nil {
		log.Println("Error in creating table:", err)
	} else {
		log.Println("Table created")
	}

}
