package main

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/DTrinadcat/crm-workers/migrations"
)

func main() {

	dsn := "postgres://admin:secret@localhost:5432/workers?sslmode=disable"
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("No connection with PostgreSQL:", err)
	}

	if err := migrations.Run(db); err != nil {
		log.Fatal("Migration Error:", err)
	}

	log.Println("BD Ready")
}
