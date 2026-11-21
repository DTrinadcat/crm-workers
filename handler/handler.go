package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/DTrinadcat/crm-workers/store"
	"github.com/google/uuid"
)

const (
	GetMainRequestMessage = "get request complete"
)

func NewRouter(s store.Store) http.Handler {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprint(w, GetMainRequestMessage)

	})

	mux.HandleFunc("GET /worker/{id}", func(w http.ResponseWriter, r *http.Request) {

		id, err := uuid.Parse(r.PathValue("id"))

		if err != nil {
			http.Error(w, "Incorrect ID (UUID)", http.StatusBadRequest)
		}

		worker, err := s.GetWorkerByID(id)

		if err != nil {
			http.Error(w, "ID was not found", http.StatusNotFound)
		}

		b, _ := json.Marshal(worker)
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	})

	mux.HandleFunc("POST /worker/{$}", func(w http.ResponseWriter, r *http.Request) {

		var worker store.Worker

		json.NewDecoder(r.Body).Decode(&worker)

		s.AddWorker(worker)

	})

	return mux

}
