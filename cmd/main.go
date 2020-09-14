package main

import (
	"RSOI/internal/pkg/middleware"
	"RSOI/internal/pkg/persona/delivery"
	"RSOI/internal/pkg/persona/repository"
	"RSOI/internal/pkg/persona/usecase"
	"context"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
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

	connection, state := os.LookupEnv("dbData")
	if !state {
		log.Fatal("connection string was not found")
	}

	conn, err := pgxpool.Connect(context.Background(), connection)
	if err != nil {
		log.Fatal("database connection not established")
	}

	pr := repository.NewPRepository(*conn)
	pu := usecase.NewPUsecase(pr)
	pd := delivery.NewPHandler(pu)

	r := mux.NewRouter()
	r.Use(middleware.InternalServerError)

	r.HandleFunc("/person/{personID}", pd.Read).Methods("GET")
	r.HandleFunc("/persons", pd.ReadAll).Methods("GET")
	r.HandleFunc("/person", pd.Create).Methods("POST")
	r.HandleFunc("/person/{personID}", pd.Update).Methods("PATCH")
	r.HandleFunc("/person/{personID}", pd.Delete).Methods("DELETE")

	srv := &http.Server{
		Handler:      r,
		Addr:         ":5000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Print("Server running at ", srv.Addr)
	log.Fatal(srv.ListenAndServe())

}

