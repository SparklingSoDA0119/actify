package rest_v1

import (
	"actify_srv/internal/db"
	rest_user_v1 "actify_srv/internal/rest/v1/user"
	"net/http"

	"github.com/gorilla/mux"
)


func RegistAllFunc(router *mux.Router, pDb *db.PostgresDB) {
	registUserFunc(router, pDb)
}


func registUserFunc(router *mux.Router, pDb *db.PostgresDB) {
	router.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		rest_user_v1.GetUserHandler(w, r, pDb)
	}).Methods("GET")

	router.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		rest_user_v1.SignUpHandler(w, r, pDb)
	}).Methods("POST")
}


