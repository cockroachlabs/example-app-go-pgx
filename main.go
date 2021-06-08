package main

import (
	"context"
	"fmt"
	"log"

	"github.com/cockroachdb/cockroach-go/v2/crdb/crdbpgx"
	"github.com/jackc/pgx/v4"
)

func transferFunds(ctx context.Context, tx pgx.Tx, from int, to int, amount int) error {
	// Read the balance.
	var fromBalance int
	if err := tx.QueryRow(ctx,
		"SELECT balance FROM accounts WHERE id = $1", from).Scan(&fromBalance); err != nil {
		return err
	}

	if fromBalance < amount {
		return fmt.Errorf("insufficient funds")
	}

	// Perform the transfer.
	if _, err := tx.Exec(ctx,
		"UPDATE accounts SET balance = balance - $1 WHERE id = $2", amount, from); err != nil {
		return err
	}
	if _, err := tx.Exec(ctx,
		"UPDATE accounts SET balance = balance + $1 WHERE id = $2", amount, to); err != nil {
		return err
	}
	return nil
}

func main() {
	config, err := pgx.ParseConfig("postgres://demo:{demo_password}@127.0.0.1:26257/bank?sslmode=require")
	if err != nil {
		log.Fatal("error configuring the database: ", err)
	}

	// Connect to the "bank" database.
	conn, err := pgx.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}
	defer conn.Close(context.Background())

	// Create the "accounts" table.
	if _, err := conn.Exec(context.Background(),
		"CREATE TABLE IF NOT EXISTS accounts (id INT PRIMARY KEY, balance INT)"); err != nil {
		log.Fatal(err)
	}

	// Insert two rows into the "accounts" table.
	if _, err := conn.Exec(context.Background(),
		"INSERT INTO accounts (id, balance) VALUES (1, 1000), (2, 250)"); err != nil {
		log.Fatal(err)
	}

	// Print out the balances.
	rows, err := conn.Query(context.Background(), "SELECT id, balance FROM accounts")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	fmt.Println("Initial balances:")
	for rows.Next() {
		var id, balance int
		if err := rows.Scan(&id, &balance); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d %d\n", id, balance)
	}

	// Run a transfer in a transaction.
	err = crdbpgx.ExecuteTx(context.Background(), conn, pgx.TxOptions{}, func(tx pgx.Tx) error {
		return transferFunds(context.Background(), tx, 1 /* from acct# */, 2 /* to acct# */, 100 /* amount */)
	})
	if err == nil {
		fmt.Println("Success")
	} else {
		log.Fatal("error: ", err)
	}
}
