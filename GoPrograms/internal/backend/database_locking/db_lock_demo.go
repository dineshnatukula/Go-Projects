package backend

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"time"

	_ "github.com/lib/pq"
)

// Repo archive storage
type Repo struct {
	db *sql.DB
}

// InitRepo initilizes repo
func InitRepo(ctx context.Context, connectionString string) (*Repo, error) {
	// DB
	postgreDb, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	err = postgreDb.Ping()
	if err != nil {
		return nil, err
	}

	postgreDb.SetMaxOpenConns(100)
	postgreDb.SetMaxIdleConns(100)
	postgreDb.SetConnMaxLifetime(5 * time.Minute)

	return &Repo{
		db: postgreDb,
	}, nil
}

func DatabaseDeadLockDemo() {

	var dbName = flag.String("db.name", "wallet", "name of the database")
	var dbHost = flag.String("db.host", "localhost", "db host")
	var dbPort = flag.String("db.port", "5432", "db port")
	var dbUser = flag.String("db.user", "wallet", "db user")
	var dbPass = flag.String("db.pass", "ppbet123", "db password")

	ctx := context.Background()

	connStr := fmt.Sprintf("dbname=%s host=%s port=%s user=%s password=%s application_name=%s sslmode=disable", *dbName, *dbHost, *dbPort, *dbUser, *dbPass, path.Base(os.Args[0]))
	db1, err := InitRepo(ctx, connStr)
	checkErr(err)
	db2, err := InitRepo(ctx, connStr)
	checkErr(err)

	go tx1(db1.db)
	time.Sleep(2 * time.Second) // Give tx1 time to acquire first lock
	go tx2(db2.db)

	// Wait to observe output
	time.Sleep(15 * time.Second)
}

func tx1(db *sql.DB) {
	tx, err := db.Begin()
	checkErr(err)

	fmt.Println("Tx1: locking RC1 (id=1)")
	_, err = tx.Exec(`SELECT * FROM accounts WHERE id = 1 FOR UPDATE`)
	checkErr(err)

	fmt.Println("Tx1: sleeping before locking RC2 (id=2)")
	time.Sleep(5 * time.Second)

	fmt.Println("Tx1: trying to lock RC2 (id=2)")
	_, err = tx.Exec(`SELECT * FROM accounts WHERE id = 2 FOR UPDATE`)
	if err != nil {
		fmt.Println("Tx1 ERROR (likely deadlock):", err)
		_ = tx.Rollback()
		return
	}

	fmt.Println("Tx1: Commit")
	_ = tx.Commit()
}

func tx2(db *sql.DB) {
	tx, err := db.Begin()
	checkErr(err)

	fmt.Println("Tx2: locking RC2 (id=2)")
	_, err = tx.Exec(`SELECT * FROM accounts WHERE id = 2 FOR UPDATE`)
	checkErr(err)

	fmt.Println("Tx2: sleeping before locking RC1 (id=1)")
	time.Sleep(5 * time.Second)

	fmt.Println("Tx2: trying to lock RC1 (id=1)")
	_, err = tx.Exec(`SELECT * FROM accounts WHERE id = 1 FOR UPDATE`)
	if err != nil {
		fmt.Println("Tx2 ERROR (likely deadlock):", err)
		_ = tx.Rollback()
		return
	}

	fmt.Println("Tx2: Commit")
	_ = tx.Commit()
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Creating tables
// CREATE TABLE accounts (
//     id INT PRIMARY KEY,
//     balance NUMERIC
// );

// INSERT INTO accounts (id, balance) VALUES (1, 1000), (2, 2000);

// Tx1: locking RC1 (id=1)
// Tx2: locking RC2 (id=2)
// Tx1: trying to lock RC2 (id=2)
// Tx2: trying to lock RC1 (id=1)
// Tx2 ERROR (likely deadlock): pq: deadlock detected
