package store

import (
	"errors"

	"github.com/google/uuid"
)

var ErrMessageNotFindID = errors.New("ID not find")
var ErrWorkerAlreadyExists = errors.New("Workers already exists")
var ErrUnidentifiedAdding = errors.New("Unidentified creation error")
var ErrUnidentifiedGetting = errors.New("Unidentified getting error")
var ErrUnidentifiedDeleting = errors.New("Unidentified deleting error")
var ErrUnidentifiedUpdating = errors.New("Unidentified updating error")

type Worker struct {
	ID                                                                    uuid.UUID
	Name, Surname, Role, Office, Number, AdditionalNumber, Email, Address string
	IsAdmin                                                               bool
	Salary                                                                int
}

type Store interface {
	GetWorkerByID(id uuid.UUID) (Worker, error)
	AddWorker(worker Worker) error
	DeleteWorker(id uuid.UUID) error
	UpdateWorker(worker Worker) error
	// GetAllWorkers
}
