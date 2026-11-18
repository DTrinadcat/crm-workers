package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"

	handler "example.com/miniprojects/crm_workers"
)

func TestHandler(t *testing.T) {

	t.Run("check main routes", func(t *testing.T) {

		serverRouter := handler.NewRouter()

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		serverRouter.ServeHTTP(response, request)

		got := response.Body.String()
		want := handler.GetMainRequestMessage

		if want != got {
			t.Errorf("want: %+v | got %+v", want, got)
		}

	})
	t.Run("check  routes", func(t *testing.T) {

		serverRouter := handler.NewRouter()

		userId := uuid.New()

		workers := []handler.Worker{{
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

		bDummy, _ := json.Marshal(workers[1])

		requestDummy := httptest.NewRequest(http.MethodPost, "/worker/", bytes.NewReader(bDummy))
		requestDummy.Header.Set("Content-Type", "application/json")
		serverRouter.ServeHTTP(httptest.NewRecorder(), requestDummy)

		b, _ := json.Marshal(workers[0])

		r := bytes.NewReader(b)

		request := httptest.NewRequest(http.MethodPost, "/worker/", r)
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()

		serverRouter.ServeHTTP(response, request)

		bGet, _ := json.Marshal(userId)

		rGet := bytes.NewReader(bGet)

		requestGet := httptest.NewRequest(http.MethodGet, "/worker/", rGet)
		requestGet.Header.Set("Content-Type", "application/json")
		responseGet := httptest.NewRecorder()

		serverRouter.ServeHTTP(responseGet, requestGet)

		var got handler.Worker

		json.Unmarshal(responseGet.Body.Bytes(), &got)

		want := workers[0]

		if want != got {
			t.Errorf("want: %+v | got %+v", want, got)
		}

	})

}
