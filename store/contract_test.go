package store_test

import (
	"context"
	"database/sql"
	"errors"
	"slices"
	"testing"

	"github.com/google/uuid"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"

	"github.com/DTrinadcat/crm-workers/migrations"
	"github.com/DTrinadcat/crm-workers/store"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func checkStoreContract(t *testing.T, s store.Store) {

	t.Run("Add Worker, Get Worker, Update Worker and Delete Worker", func(t *testing.T) {

		var worker = store.Worker{
			ID:               uuid.New(),
			Name:             "Martin",
			Surname:          "Eden",
			Role:             "Writer",
			Office:           "USA, California, Oakland, Niche Avenue, 34",
			Number:           "+12220009542",
			AdditionalNumber: "+13944419844",
			Email:            "martin.eden@example.com",
			Address:          "USA, California, Oakland, Good street, 30, 5",
			IsAdmin:          false,
			Salary:           200_000,
		}

		var workerUpdated = store.Worker{
			ID:               worker.ID,
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

		// Add
		if err := s.AddWorker(worker); err != nil {

			t.Fatal("Adding error:", err.Error())

		}

		// Get
		receivedWorker, err := s.GetWorkerByID(worker.ID)

		if err != nil {
			t.Fatal("Getting error:", err.Error())
		}

		if worker != receivedWorker {
			t.Errorf("want: %+v | got %+v", worker, receivedWorker)
		}

		// Update

		if err := s.UpdateWorker(workerUpdated); err != nil {
			t.Fatal("Updating error:", err.Error())
		}

		// Get Updated

		receivedWorker, err = s.GetWorkerByID(worker.ID)

		if err != nil {
			t.Fatal("Getting error:", err.Error())
		}

		if workerUpdated != receivedWorker {
			t.Fatalf("want: %+v | got %+v", workerUpdated, receivedWorker)
		}

		// Delete

		if err := s.DeleteWorker(workerUpdated.ID); err != nil {
			t.Fatal("Deleting error:", err.Error())
		}

		// Get Deleted

		if _, err = s.GetWorkerByID(worker.ID); err != nil {

			if !errors.Is(err, store.ErrMessageNotFindID) {
				t.Fatal("Something wrong with deletion", err.Error())
			}

		} else {
			t.Fatal("Data wasn't deleted")
		}

	})

	t.Run("Get All Workers", func(t *testing.T) {

		var worker = store.Worker{
			ID:               uuid.New(),
			Name:             "Martin",
			Surname:          "Eden",
			Role:             "Writer",
			Office:           "USA, California, Oakland, Niche Avenue, 34",
			Number:           "+12220009542",
			AdditionalNumber: "+13944419844",
			Email:            "martin.eden@example.com",
			Address:          "USA, California, Oakland, Good street, 30, 5",
			IsAdmin:          false,
			Salary:           200_000,
		}

		var worker2 = store.Worker{
			ID:               uuid.New(),
			Name:             "Evgeniy",
			Surname:          "Onegin",
			Role:             "dandy",
			Office:           "Moscow, Arbat, 3",
			Number:           "+79998886665",
			AdditionalNumber: "",
			Email:            "evgeniy.onegin@example.com",
			Address:          "Moscow, uncle house",
			IsAdmin:          false,
			Salary:           200_000_000,
		}

		// Add
		if err := s.AddWorker(worker); err != nil {

			t.Fatal("Adding error:", err.Error())

		}

		// Add
		if err := s.AddWorker(worker2); err != nil {

			t.Fatal("Adding error:", err.Error())

		}

		// Get All
		workers := []store.Worker{worker, worker2}

		receivedWorkers, err := s.GetAllWorkers()

		if err != nil {

			t.Fatal("Get All error:", err.Error())

		}

		if !slices.Equal(workers, receivedWorkers) {

			t.Errorf("want: %+v | got %+v", workers, receivedWorkers)

		}

	})

}

func TestMemoryStore(t *testing.T) {

	s := store.NewMemoryStore()

	checkStoreContract(t, s)

}

func TestPSQLStore(t *testing.T) {

	ctx := context.Background()

	pgContainer, err := postgres.Run(ctx, "postgres:17-alpine",

		postgres.WithDatabase("workers_test"),
		postgres.WithUsername("admin"),
		postgres.WithPassword("secret"),
		postgres.BasicWaitStrategies(),
	)

	if err != nil {
		t.Fatal("Docker is probably not running")
	}

	t.Cleanup(func() {

		if err := testcontainers.TerminateContainer(pgContainer); err != nil {

			t.Logf("cannot stop container: %v", err)

		}

	})

	dsn, err := pgContainer.ConnectionString(ctx, "sslmode=disable")

	if err != nil {
		t.Error("connection string error")
		return
	}

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		db.Close()
	})

	if err := db.Ping(); err != nil {
		t.Fatal("No connection with PostgreSQL:", err)
	}

	if err := migrations.Run(db); err != nil {
		t.Fatal("Migration Error:", err)
	}

	s := store.NewPSQLStore(db)

	checkStoreContract(t, s)

}
