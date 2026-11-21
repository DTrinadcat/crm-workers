package store

import (
	"errors"

	"github.com/google/uuid"
)

type Worker struct {
	ID                                                                    uuid.UUID
	Name, Surname, Role, Office, Number, AdditionalNumber, Email, Address string
	IsAdmin                                                               bool
	Salary                                                                int
}

type Store interface {
	GetWorkerByID(id uuid.UUID) (Worker, error)
	AddWorker(worker Worker)
}

type MemoryStore []Worker

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{}
}

func (ms *MemoryStore) GetWorkerByID(id uuid.UUID) (Worker, error) {
	for _, worker := range *ms {
		if id == worker.ID {
			return worker, nil
		}
	}
	return Worker{}, errors.New("Not find that id")
}

func (ms *MemoryStore) AddWorker(worker Worker) {
	*ms = append(*ms, worker)
}
