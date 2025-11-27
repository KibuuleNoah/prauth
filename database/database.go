package database

import (
	"context"
	"log"
	"prauth/services"

	"github.com/jackc/pgx/v5"
)


func InitDB() (*pgx.Conn,context.Context) {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, services.GetDBURL())
	if err != nil {
		log.Fatal("Unable to connect:", err)
	}

	if err = createUserTable(conn, ctx); err != nil{
		log.Fatal(err)
	}
	
	return conn, ctx
}

func createUserTable(conn *pgx.Conn, ctx context.Context) error {
	_, err := conn.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL,
			email TEXT UNIQUE NOT NULL,
			pwdhash TEXT NOT NULL
		)
	`)
	return err
}

