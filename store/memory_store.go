package store

import (
	"slices"

	"github.com/google/uuid"
)

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
	return Worker{}, ErrMessageNotFindID
}

func (ms *MemoryStore) AddWorker(worker Worker) error {
	*ms = append(*ms, worker)

	return nil
}

func (ms *MemoryStore) DeleteWorker(id uuid.UUID) error {

	var isDeleted bool
	*ms = slices.DeleteFunc(*ms, func(worker Worker) bool {
		if worker.ID == id {
			isDeleted = true
			return true
		}

		return false
	})

	if isDeleted {
		return nil
	}

	return ErrMessageNotFindID

}

func (ms *MemoryStore) UpdateWorker(worker Worker) error {

	for i, el := range *ms {

		if el.ID == worker.ID {
			(*ms)[i] = worker
			return nil
		}

	}

	return ErrMessageNotFindID
}

func (ms *MemoryStore) GetAllWorkers() ([]Worker, error) {
	return *ms, nil
}
