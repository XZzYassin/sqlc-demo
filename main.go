package main

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"test.mumen.io/repo"
)

func run() error {
	ctx := context.Background()

	db, err := sql.Open("postgres", "user=postgres dbname=test-db sslmode=disable")
	if err != nil {
		return err
	}

	queries := repo.New(db)

	// list user with id 1
	user, err := queries.GetUser(ctx, 1)
	if err != nil {
		return err
	}
	log.Println(user)

	// list all users
	users, err := queries.GetUsers(ctx)
	if err != nil {
		return err
	}
	log.Println(users)

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
