package rest_v1

import (
	"actify_srv/internal/db"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)


func RegistAllFunc(router *mux.Router, pDb *db.PostgresDB) {
	router.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		getUserHandler(w, r, pDb)
	}).Methods("Get")
}


func getUserHandler(w http.ResponseWriter, r *http.Request, pDb *db.PostgresDB) {
	data := map[string]string {
		"message": "hello world",
		"status": "sucess",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

