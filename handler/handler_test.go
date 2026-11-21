package handler_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DTrinadcat/crm-workers/handler"
	"github.com/DTrinadcat/crm-workers/store"
	"github.com/google/uuid"
)

func routeData(data any, response *httptest.ResponseRecorder, address string, mux http.Handler, method string) {
	// TODO add an error handler
	b, _ := json.Marshal(data)
	request := httptest.NewRequest(method, address, bytes.NewReader(b))
	request.Header.Set("Content-Type", "application/json")
	mux.ServeHTTP(response, request)
}

func routeDataPathValue(response *httptest.ResponseRecorder, address string, mux http.Handler, method string) {

	request := httptest.NewRequest(method, address, nil)
	mux.ServeHTTP(response, request)

}

func TestHandler(t *testing.T) {

	t.Run("check main routes (GET /)", func(t *testing.T) {

		serverRouter := handler.NewRouter(store.NewMemoryStore())

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		serverRouter.ServeHTTP(response, request)

		got := response.Body.String()
		want := handler.GetMainRequestMessage

		if want != got {
			t.Errorf("want: %+v | got %+v", want, got)
		}

	})
	t.Run("check get worker route (POST + GET /worker/)", func(t *testing.T) {

		s := store.NewMemoryStore()

		serverRouter := handler.NewRouter(s)

		userId := uuid.New()

		workers := []store.Worker{{
			ID:               userId,
			Name:             "John",
			Surname:          "Smith",
			Role:             "GO Developer",
			Office:           "New York, Fifth Avenue, 3",
			Number:           "+13131311313",
			AdditionalNumber: "+13009403999",
			Email:            "john.smith@example.com",
			Address:          "New York, Van Buren, 13, 90",
			IsAdmin:          true,
			Salary:           100_000},
			{
				ID:               uuid.New(),
				Name:             "Elena",
				Surname:          "Smirnova",
				Role:             "Frontend Lead",
				Office:           "Moscow, Tverskaya, 12",
				Number:           "+79001112233",
				AdditionalNumber: "+79004445566",
				Email:            "elena.smirnova@example.com",
				Address:          "Moscow, Arbat 24, 15",
				IsAdmin:          false,
				Salary:           150_000,
			},
		}

		// Dummy Post
		routeData(workers[1], httptest.NewRecorder(), "/worker/", serverRouter, http.MethodPost)

		// Target Post
		routeData(workers[0], httptest.NewRecorder(), "/worker/", serverRouter, http.MethodPost)

		// Target Get
		response := httptest.NewRecorder()
		routeDataPathValue(response, fmt.Sprintf("/worker/{%s}", userId), serverRouter, http.MethodGet)

		var got store.Worker

		json.Unmarshal(response.Body.Bytes(), &got)

		want := workers[0]

		if want != got {
			t.Errorf("want: %+v | got %+v", want, got)
		}

	})

}
