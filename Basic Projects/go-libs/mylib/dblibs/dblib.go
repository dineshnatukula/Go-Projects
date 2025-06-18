// ~/go-libs/mylib/mylib.go
package dblibs

import (
	"database/sql"
	"fmt"
	"log"
)

func DBDemo() {
	fmt.Println("DB Init Started....")

	fmt.Println("DB Init Ended......")
}

type Datadb struct {
	DB *sql.DB
}

func DBInit(dbHost, dbPort, dbDatabaseName, dbUserName, dbPassword string) *sql.DB {
	connStr := "user=postgres password=root dbname=College sslmode=disable"

	ddb := new(Datadb)

	// Open a database connection
	DB, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	ddb.DB = DB

	// Example of setting a maximum connection pool size
	// ddb.DB.SetMaxOpenConns(10)

	defer ddb.DB.Close()
	return nil
}
