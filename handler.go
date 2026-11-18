package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

const (
	GetMainRequestMessage = "get request complete"
)

type Worker struct {
	ID                                                                    uuid.UUID
	Name, Surname, Role, Office, Number, AdditionalNumber, Email, Address string
	IsAdmin                                                               bool
	Salary                                                                int
}

func NewRouter() http.Handler {

	var workers []Worker

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprint(w, GetMainRequestMessage)

	})

	mux.HandleFunc("GET /worker/{$}", func(w http.ResponseWriter, r *http.Request) {

		var id uuid.UUID

		json.NewDecoder(r.Body).Decode(&id)

		for _, worker := range workers {

			if worker.ID == id {
				b, _ := json.Marshal(worker)
				w.Write(b)
				return
			}

		}

		fmt.Print("Worker not found")

	})

	mux.HandleFunc("POST /worker/{$}", func(w http.ResponseWriter, r *http.Request) {

		var worker Worker

		json.NewDecoder(r.Body).Decode(&worker)

		workers = append(workers, worker)

	})

	return mux

}
