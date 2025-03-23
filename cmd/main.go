package main

import (
	"fmt"
	"go-project/pkg/api"
	"go-project/pkg/repository"
	"log"
	"net/http"

	"github.com/rs/cors"

	"github.com/gorilla/mux"
)

const connStr = "postgres://postgres:1258@localhost:5432/siteDaewoo"

func main() {
	router := mux.NewRouter()
	db, err := repository.New(connStr)
	if err != nil {
		log.Fatal(err.Error())
	}

	api := api.New(router, db)
	api.Handle()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
	})

	handler := c.Handler(router)

	fmt.Println("Сервер запущен на http://localhost:8080")
	log.Fatal(http.ListenAndServe("localhost:8080", handler))
}
