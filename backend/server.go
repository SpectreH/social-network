package main

import (
	"database/sql"
	"log"
	"net/http"
	"social-network/internal/database/sqlite"
	"social-network/internal/handlers"
	"social-network/internal/routes"
)

func main() {
	database, _ := sql.Open("sqlite3", "./db/network.db")
	defer database.Close()

	err := sqlite.UseMigrations(database)
	if err != nil {
		log.Fatal(err)
	}

	handlersRepo := handlers.SetNewRepo(database)
	handlers.SetNewHandlers(handlersRepo)

	srv := &http.Server{
		Addr:    ":4000",
		Handler: routes.SetRoutes(),
	}

	log.Println("Starting application on port " + srv.Addr)
	if srv.ListenAndServe() != nil {
		log.Fatalf("%v - Internal Server Error", http.StatusInternalServerError)
	}
}
