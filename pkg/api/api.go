// package api

// import (
// 	"go-project/pkg/handlers"
// 	"go-project/pkg/repository"
// 	"net/http"

// 	"github.com/gorilla/mux"
// )

// type api struct {
// 	r  *mux.Router
// 	db *repository.PGRepo
// }

// func New(router *mux.Router, db *repository.PGRepo) *api {
// 	return &api{r: router, db: db}
// }

// func (api *api) Handle() {
// 	api.r.HandleFunc("/api/hello", handlers.Hello).Methods(http.MethodGet).Queries("name", "{name}")
// }

package api

import (
	"go-project/pkg/handlers"
	"go-project/pkg/repository"
	"net/http"

	"github.com/gorilla/mux"
)

type api struct {
	r  *mux.Router
	db *repository.PGRepo
}

func New(router *mux.Router, db *repository.PGRepo) *api {
	return &api{r: router, db: db}
}

func (api *api) Handle() {
	api.r.HandleFunc("/contacts", handlers.PostContact(api.db)).Methods("POST")

}

func (api *api) ListenAndServe(addr string) error {
	return http.ListenAndServe(addr, api.r)
}
