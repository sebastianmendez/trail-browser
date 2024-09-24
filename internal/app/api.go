package app

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sebastianmendez/trail-browser/internal/app/controller"
)

type API struct {
	BasePath string
	Addr     string
	server   *http.Server
}

// boliler plate setup to handle http requests using golang standard library
func Start(basePath, address string) API {
	a := API{
		BasePath: basePath,
		Addr:     address,
	}

	a.server = &http.Server{
		Addr: a.Addr,
		Handler: func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				next.ServeHTTP(w, r)
			})
		}(a.createRouter()),
	}

	return a
}
func (a API) createRouter() *mux.Router {
	router := mux.NewRouter().PathPrefix(a.BasePath).Subrouter()
	router.HandleFunc("/trails", controller.HandleList).Methods("GET")
	return router
}

func (a API) ListenAndServe(errChan chan error) {
	err := a.server.ListenAndServe()
	if err != nil {
		errChan <- errors.New("httpAPI was shutdown or closed:" + err.Error())
	}
}
