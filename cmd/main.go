package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/google/uuid"
	_ "github.com/jackc/pgx/v5/stdlib"

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

	psqlDB.DeleteWorker(uuid.MustParse("a3d97519-4805-4093-af6b-403888fd0f4f"))

	worker, err := psqlDB.GetWorkerByID(uuid.MustParse("a3d97519-4805-4093-af6b-403888fd0f4f"))
	fmt.Println(worker)

}
