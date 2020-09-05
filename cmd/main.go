package main

import (
	"RSOI/internal/pkg/persona/delivery"
	"RSOI/internal/pkg/persona/repository"
	"RSOI/internal/pkg/persona/usecase"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {

	pr := repository.NewPRepository()
	pu := usecase.NewPUsecase(pr)
	pd := delivery.NewPHandler(pu)

	r := mux.NewRouter()

	r.HandleFunc("/person/{personID}", pd.Read).Methods("GET")
	r.HandleFunc("/person", pd.ReadAll).Methods("GET")
	r.HandleFunc("/person", pd.Create).Methods("POST")
	r.HandleFunc("/person/{personID}", pd.Update).Methods("PATCH")
	r.HandleFunc("/person/{personID}", pd.Delete).Methods("DELETE")

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

}
