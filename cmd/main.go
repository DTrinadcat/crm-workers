package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/google/uuid"
	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/DTrinadcat/crm-workers/handler"
	"github.com/DTrinadcat/crm-workers/migrations"
	"github.com/DTrinadcat/crm-workers/store"
)

// var workerEden = store.Worker{
// 	ID:               uuid.New(),
// 	Name:             "Martin",
// 	Surname:          "Eden",
// 	Role:             "Writer",
// 	Office:           "USA, California, Oakland, Niche Avenue, 34",
// 	Number:           "+12220009542",
// 	AdditionalNumber: "+13944419844",
// 	Email:            "martin.eden@example.com",
// 	Address:          "USA, California, Oakland, Good street, 30, 5",
// 	IsAdmin:          false,
// 	Salary:           200_000,
// }

var workerEdenUpdated = store.Worker{
	ID:               uuid.MustParse("57347eab-10cb-431b-b685-0fdcdc61bd63"),
	Name:             "Martin",
	Surname:          "Eden",
	Role:             "Ex-Writer",
	Office:           "USA, California, Oakland, Niche Avenue, 34",
	Number:           "+12220009542",
	AdditionalNumber: "+13944419844",
	Email:            "martin.eden@example.com",
	Address:          "sea",
	IsAdmin:          false,
	Salary:           100_000_000,
}

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

	psqlDB := store.NewPSQLStore(db)

	mux := handler.NewRouter(psqlDB)

	http.ListenAndServe(":8080", mux)

}
