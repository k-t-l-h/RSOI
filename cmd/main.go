package main

import (
	"RSOI/internal/pkg/persona/delivery"
	"RSOI/internal/pkg/persona/repository"
	"RSOI/internal/pkg/persona/usecase"
	"context"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/pgxpool"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"time"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	connection, _ := os.LookupEnv("dbData")

	conn, err := pgxpool.Connect(context.Background(), connection)
	if err != nil {
		log.Fatal("database error")
	}

	pr := repository.NewPRepository(*conn)
	pu := usecase.NewPUsecase(pr)
	pd := delivery.NewPHandler(pu)

	r := mux.NewRouter()

	r.HandleFunc("/person/{personID:[0-9]+}", pd.Read).Methods("GET")
	r.HandleFunc("/persons", pd.ReadAll).Methods("GET")
	r.HandleFunc("/person", pd.Create).Methods("POST")
	r.HandleFunc("/person/{personID}", pd.Update).Methods("PATCH")
	r.HandleFunc("/person/{personID}", pd.Delete).Methods("DELETE")

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:5000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

}
