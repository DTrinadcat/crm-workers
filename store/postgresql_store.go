package store

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
)

type PSQLStore struct {
	db *sql.DB
}

func NewPSQLStore(db *sql.DB) *PSQLStore {
	return &PSQLStore{db: db}
}

func (psqlS *PSQLStore) AddWorker(worker Worker) error {

	_, err := psqlS.db.Exec("insert into workers (id, name, surname, role, office, number, additional_number, email, address, is_admin, salary) values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)", worker.ID, worker.Name, worker.Surname, worker.Role, worker.Office, worker.Number, worker.AdditionalNumber, worker.Email, worker.Address, worker.IsAdmin, worker.Salary)

	if err != nil {

		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return ErrWorkerAlreadyExists
		}

		return ErrUnidentifiedAdding

	}

	return nil

}

func (psqlS *PSQLStore) DeleteWorker(id uuid.UUID) error {
	res, err := psqlS.db.Exec("delete from workers where id = $1", id)

	if err != nil {
		return ErrUnidentifiedDeleting
	}

	rows, err := res.RowsAffected()

	if err != nil {
		return ErrUnidentifiedDeleting
	}

	if rows == 0 {
		return ErrMessageNotFindID
	}

	return nil
}

func (psqlS *PSQLStore) GetWorkerByID(id uuid.UUID) (Worker, error) {

	result := psqlS.db.QueryRow("select * from workers where id = $1", id)

	var worker Worker

	if err := result.Scan(&worker.ID, &worker.Name, &worker.Surname, &worker.Role, &worker.Office, &worker.Number, &worker.AdditionalNumber, &worker.Email, &worker.Address, &worker.IsAdmin, &worker.Salary); err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return worker, ErrMessageNotFindID
		}

		return worker, ErrUnidentifiedGetting

	}

	return worker, nil

}

func (psqlS *PSQLStore) UpdateWorker(worker Worker) error {

	result, err := psqlS.db.Exec("update workers set (name, surname, role, office, number, additional_number, email, address, is_admin, salary) = ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10) where id = $11", worker.Name, worker.Surname, worker.Role, worker.Office, worker.Number, worker.AdditionalNumber, worker.Email, worker.Address, worker.IsAdmin, worker.Salary, worker.ID)

	if err != nil {

		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return ErrWorkerAlreadyExists
		}

		return ErrUnidentifiedUpdating
	}

	rows, err := result.RowsAffected()

	if err != nil {
		return ErrUnidentifiedUpdating
	}

	if rows == 0 {
		return ErrMessageNotFindID
	}

	return nil
}

func (psqlS *PSQLStore) GetAllWorkers() ([]Worker, error) {

	res, err := psqlS.db.Query("select * from workers")

	if err != nil {
		return nil, ErrUnidentifiedGetting
	}

	defer res.Close()

	var workers []Worker

	for res.Next() {

		var worker Worker

		if err := res.Scan(&worker.ID, &worker.Name, &worker.Surname, &worker.Role, &worker.Office, &worker.Number, &worker.AdditionalNumber, &worker.Email, &worker.Address, &worker.IsAdmin, &worker.Salary); err != nil {

			return nil, ErrUnidentifiedGetting

		}

		workers = append(workers, worker)

	}

	if err := res.Err(); err != nil {
		return nil, ErrUnidentifiedGetting
	}

	return workers, nil

}
