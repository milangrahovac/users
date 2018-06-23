package web

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/gorilla/mux"

	"github.com/milangrahovac/users/client/users"
)

func NewRouter(log *logrus.Logger) http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/users", createUser(log)).Methods(http.MethodPost)
	r.HandleFunc("api/v1/login", login).Methods(http.MethodPost)

	return r
}

func createUser(log *logrus.Logger) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		user := &users.RequestCreateUser{}

		err := json.NewDecoder(r.Body).Decode(user)
		if err != nil {
			res := users.Response{
				Error: &users.ResponseError{
					Code:        users.InvalidBody,
					Description: "Couldn't parse JSON body.",
				},
			}

			w.Header().Set("Content-Type", "application/json; charset=utf=8")
			w.WriteHeader(http.StatusBadRequest)

			err := json.NewEncoder(w).Encode(res)
			if err != nil {
				log.Errorf("Couldn't encode error message `%s`", err)
			}
			return
		}
		log.Infof("%+v", user)
	}
}

func login(w http.ResponseWriter, r *http.Request) {

}
