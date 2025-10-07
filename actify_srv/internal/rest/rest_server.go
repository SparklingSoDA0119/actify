package rest

import (
	"fmt"
	"net/http"
	"time"

	"actify_srv/internal/db"
	rest_v1 "actify_srv/internal/rest/v1"

	"github.com/gorilla/mux"
)


type RestServer struct {
	server 		*http.Server
	isOpened	bool
	api		*mux.Router
	pDb         *db.PostgresDB
}


func NewRestServer(pDb *db.PostgresDB) *RestServer {
	return &RestServer{
		isOpened: false,
		pDb: pDb,
	}
}


func (srv *RestServer) Initialize(addr string) {
	if srv.isOpened {
		return 
	}

	r := mux.NewRouter()

	srv.api = r.PathPrefix("/actify/data").Subrouter()
	rest_v1.RegistAllFunc(srv.api.PathPrefix("/v1").Subrouter(), srv.pDb)

	srv.server = &http.Server{
		Handler: r,
		Addr: addr,
		WriteTimeout: 10 * time.Second,
		ReadTimeout: 10 * time.Second,
	}

	srv.isOpened = true
}


func (srv *RestServer) Listen() error {
	if !srv.isOpened {
		return fmt.Errorf("server not initialized")
	}

	srv.server.ListenAndServe()
	return nil
}