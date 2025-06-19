package backend

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v5"
)

const dsn = "postgres://postgres:password@localhost:5432/testdb" // change as needed

func DatabaseDeadLockRetryMech() {
	ctx := context.Background()

	conn1, err := pgx.Connect(ctx, dsn)
	checkErr(err)
	defer conn1.Close(ctx)

	conn2, err := pgx.Connect(ctx, dsn)
	checkErr(err)
	defer conn2.Close(ctx)

	// Run transactions in parallel
	go runWithRetry(ctx, "Tx3", conn1, tx3)
	time.Sleep(2 * time.Second) // Let Tx3 lock RC1 first
	go runWithRetry(ctx, "Tx4", conn2, tx4)

	time.Sleep(20 * time.Second) // Wait long enough for transactions to complete
}

func runWithRetry(ctx context.Context, name string, conn *pgx.Conn, fn func(pgx.Tx) error) {
	const maxRetries = 3

	for i := 1; i <= maxRetries; i++ {
		fmt.Printf("%s: attempt %d\n", name, i)

		err := pgx.BeginFunc(ctx, conn, fn)
		if err != nil {
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) && pgErr.Code == "40P01" {
				fmt.Printf("%s: deadlock detected, retrying...\n", name)
				time.Sleep(2 * time.Second)
				continue
			} else {
				fmt.Printf("%s: failed with error: %v\n", name, err)
				return
			}
		} else {
			fmt.Printf("%s: transaction succeeded\n", name)
			return
		}
	}

	fmt.Printf("%s: failed after %d retries\n", name, maxRetries)
}

func tx3(tx pgx.Tx) error {
	ctx := context.Background()
	fmt.Println("Tx3: locking id=1")
	_, err := tx.Exec(ctx, `SELECT * FROM accounts WHERE id = 1 FOR UPDATE`)
	if err != nil {
		return err
	}

	time.Sleep(5 * time.Second) // Simulate delay

	fmt.Println("Tx3: trying to lock id=2")
	_, err = tx.Exec(ctx, `SELECT * FROM accounts WHERE id = 2 FOR UPDATE`)
	if err != nil {
		return err
	}

	fmt.Println("Tx3: done")
	return nil
}

func tx4(tx pgx.Tx) error {
	ctx := context.Background()
	fmt.Println("Tx4: locking id=2")
	_, err := tx.Exec(ctx, `SELECT * FROM accounts WHERE id = 2 FOR UPDATE`)
	if err != nil {
		return err
	}

	time.Sleep(5 * time.Second) // Simulate delay

	fmt.Println("Tx4: trying to lock id=1")
	_, err = tx.Exec(ctx, `SELECT * FROM accounts WHERE id = 1 FOR UPDATE`)
	if err != nil {
		return err
	}

	fmt.Println("Tx4: done")
	return nil
}

// func checkErr(err error) {
// 	if err != nil {
// 		log.Fatalf("Error: %v", err)
// 	}
// }

// Tx1: attempt 1
// Tx1: locking id=1
// Tx2: attempt 1
// Tx2: locking id=2
// Tx1: trying to lock id=2
// Tx2: trying to lock id=1
// Tx2: deadlock detected, retrying...
// Tx2: attempt 2
// Tx2: locking id=2
// Tx2: trying to lock id=1
// Tx2: transaction succeeded
// Tx1: done
// Tx1: transaction succeeded
