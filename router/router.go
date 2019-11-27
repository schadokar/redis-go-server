package router

import (
	"redis-server/storages"

	"github.com/gorilla/mux"
)

// Router is exported and used in main.go

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/set", storages.SetKey).Methods("POST")
	router.HandleFunc("/get", storages.Get).Methods("GET")
	router.HandleFunc("/get/{key}", storages.GetKey).Methods("GET")

	return router

}
