package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB
var err error

func InitDB() {
	connStr := "user=postgres password=postgres dbname=pokemon sslmode=disable"
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	pingErr := DB.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
	// age := 21
	// rows, err := db.Query("SELECT name FROM users WHERE age = $1", age)
}
