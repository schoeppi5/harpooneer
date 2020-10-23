package api

import (
	"github.com/gorilla/mux"
)

func NewGorilla() *mux.Router {
	router := mux.NewRouter()

	return router
}
