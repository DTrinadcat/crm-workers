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
	ID               uuid.UUID `json:"id"`
	Name             string    `json:"name"`
	Surname          string    `json:"surname"`
	Role             string    `json:"role"`
	Office           string    `json:"office"`
	Number           string    `json:"number"`
	AdditionalNumber string    `json:"additional_number"`
	Email            string    `json:"email"`
	Address          string    `json:"address"`
	IsAdmin          bool      `json:"is_admin"`
	Salary           int       `json:"salary"`
}

type Store interface {
	GetWorkerByID(id uuid.UUID) (Worker, error)
	AddWorker(worker Worker) error
	DeleteWorker(id uuid.UUID) error
	UpdateWorker(worker Worker) error
	GetAllWorkers() ([]Worker, error)
}
