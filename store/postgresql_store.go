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
	_, err := psqlS.db.Exec("delete from workers where id = $1", id)

	// Правильную ли ошибку, я здесь проверяю?
	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {

			return ErrMessageNotFindID

		}

		return ErrUnidentifiedDeleting

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

		return worker, ErrUnidentifiedAdding

	}

	return worker, nil

}
