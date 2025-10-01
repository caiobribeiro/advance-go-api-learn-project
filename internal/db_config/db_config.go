package dbconfig

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func ConnectDb(databaseUrl string) *sql.DB {
	db, err := sql.Open("postgres", databaseUrl)
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to the database!")

	return db
}
