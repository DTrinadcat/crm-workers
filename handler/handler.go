package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/DTrinadcat/crm-workers/store"
	"github.com/google/uuid"
)

func NewRouter(s store.Store) http.Handler {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /workers", func(w http.ResponseWriter, r *http.Request) {

		workers, err := s.GetAllWorkers()

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		b, err := json.Marshal(workers)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(b)

	})

	mux.HandleFunc("GET /worker/{id}", func(w http.ResponseWriter, r *http.Request) {

		id, err := uuid.Parse(r.PathValue("id"))

		if err != nil {
			http.Error(w, "Incorrect ID (UUID)", http.StatusBadRequest)
			return
		}

		worker, err := s.GetWorkerByID(id)

		if err != nil {

			if errors.Is(err, store.ErrMessageNotFindID) {
				http.Error(w, "ID was not found", http.StatusNotFound)
				return
			}

			http.Error(w, err.Error(), http.StatusInternalServerError)
			return

		}

		b, _ := json.Marshal(worker)
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	})

	mux.HandleFunc("PUT /worker", func(w http.ResponseWriter, r *http.Request) {

		var worker store.Worker

		decoder := json.NewDecoder(r.Body)

		if err := decoder.Decode(&worker); err != nil {
			http.Error(w, "Incorrect Worker instance", http.StatusBadRequest)
			return
		}

		if err := s.UpdateWorker(worker); err != nil {

			if errors.Is(err, store.ErrWorkerAlreadyExists) {
				http.Error(w, store.ErrWorkerAlreadyExists.Error(), http.StatusConflict)
				return
			} else if errors.Is(err, store.ErrMessageNotFindID) {
				http.Error(w, store.ErrMessageNotFindID.Error(), http.StatusNotFound)
				return
			}

			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)

	})

	mux.HandleFunc("POST /worker", func(w http.ResponseWriter, r *http.Request) {

		var worker store.Worker

		decoder := json.NewDecoder(r.Body)

		if err := decoder.Decode(&worker); err != nil {
			http.Error(w, "Incorrect Worker instance", http.StatusBadRequest)
			return
		}

		if err := s.AddWorker(worker); err != nil {

			if errors.Is(err, store.ErrWorkerAlreadyExists) {
				http.Error(w, store.ErrWorkerAlreadyExists.Error(), http.StatusConflict)
				return
			}

			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)

	})

	mux.HandleFunc("DELETE /worker/{id}", func(w http.ResponseWriter, r *http.Request) {

		id, err := uuid.Parse(r.PathValue("id"))

		if err != nil {
			http.Error(w, "Incorrect ID (UUID)", http.StatusBadRequest)
			return
		}

		if err := s.DeleteWorker(id); err != nil {

			if errors.Is(err, store.ErrMessageNotFindID) {
				http.Error(w, store.ErrMessageNotFindID.Error(), http.StatusNotFound)
			} else {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}

			return
		}

		w.WriteHeader(http.StatusNoContent)

	})

	return mux

}
