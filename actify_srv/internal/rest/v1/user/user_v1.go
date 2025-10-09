package rest_user_v1

import (
	"actify_srv/internal/common"
	"actify_srv/internal/db"
	"encoding/json"

	"net/http"
)

func GetUserHandler(w http.ResponseWriter, r *http.Request, pDb *db.PostgresDB) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	resp := common.RestJsonResp{
		Code: 200,
		Message: "OK",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}


func SignUpHandler(w http.ResponseWriter, r *http.Request, pDb *db.PostgresDB) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	query := r.URL.Query()
	signUpType := query.Get("type")

	if signUpType == "" || signUpType == "actify" {
		SignUpNormalHandler(w, r, pDb)
	} else {
		http.Error(w, "Not support other sign up process", http.StatusNotImplemented)
	}
}


func SignUpNormalHandler(w http.ResponseWriter, r *http.Request, pDb *db.PostgresDB) {

}