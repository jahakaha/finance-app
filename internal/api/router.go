package api

import (
	"net/http"

	"github.com/gorilla/mux"

	v1 "finance-app/internal/api/v1"
)

func NewRouter() (http.Handler, error) {
	router := mux.NewRouter()

	router.HandleFunc("/version", v1.VersionHandler)

	return router, nil
}
