package web

import (
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	reform "gopkg.in/reform.v1"

	"github.com/milangrahovac/users/client/users"
	"github.com/milangrahovac/users/internal/models"
)

func NewRouter(log *logrus.Logger, db *reform.DB) http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/users", createUser(log, db)).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/login", login).Methods(http.MethodPost)
	return r
}

func createUser(log *logrus.Logger, db *reform.DB) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		user := &users.RequestCreateUser{}
		err := json.NewDecoder(r.Body).Decode(user)
		if err != nil {
			makeErrorResponse(w, users.InvalidBody, "Couldn't parse JSON body.")
			return
		}

		// log.Debugf("%+v", user)

		if !checkName(user.Name) {
			makeErrorResponse(w, users.InvalidBody, "Name can only contain latin letters.")
			return
		}

		if !checkEmail(user.Email) {
			makeErrorResponse(w, users.InvalidBody, "Email is not correct.")
			return
		}

		if !checkPassword(user.Password) {
			makeErrorResponse(w, users.InvalidBody, "Password should contain 8-64 characters, mimimum one uppercase, minimum one lowercase and minimum one numer.")
			return
		}

		dbUser := &models.User{
			Name:     user.Name,
			Email:    user.Email,
			Password: user.Password,
		}
		err = db.Save(dbUser)
		if err != nil {
			log.Errorf("Couldn't save user in the DB: `%s`", err)
		}
	}

}

func login(w http.ResponseWriter, r *http.Request) {

}

func checkUserData(u *users.RequestCreateUser) bool {
	if len(u.Name) > 0 && len(u.Email) > 0 && len(u.Password) > 0 {
		if checkName(u.Name) && checkEmail(u.Email) && checkPassword(u.Password) {
			return true
		}
	}
	return false
}

func checkName(s string) bool {
	re := regexp.MustCompile("^[a-zA-Z]+$")
	if len(re.FindString(s)) > 0 {
		return true
	}
	return false
}

func checkPassword(s string) bool {

	re := regexp.MustCompile("^[a-zA-Z0-9]{8,64}$")
	up := regexp.MustCompile("[A-Z]+")
	lo := regexp.MustCompile("[a-z]+")
	d := regexp.MustCompile("[0-9]+")

	if len(re.FindString(s)) > 0 {
		if len(up.FindString(s)) > 0 && len(lo.FindString(s)) > 0 && len(d.FindString(s)) > 0 {
			return true
		}
	}
	return false
}

func checkEmail(s string) bool {
	re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if len(re.FindString(s)) > 0 {
		return true
	}
	return false
}

func makeErrorResponse(w http.ResponseWriter, code, description string) {
	res := users.Response{
		Error: &users.ResponseError{
			Code:        code,
			Description: description,
		},
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusBadRequest)

	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		// log.Errorf("Couldn't encode message: `%s`", err)
	}
}
