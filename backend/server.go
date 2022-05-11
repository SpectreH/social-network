package main

import (
	"database/sql"
	"log"
	"net/http"
	"social-network/internal/config"
	"social-network/internal/database/sqlite"
	"social-network/internal/handlers"
	middleware "social-network/internal/middleware"
	"social-network/internal/routes"
)

func main() {
	database, _ := sql.Open("sqlite3", config.DATABASE_PATH)
	defer database.Close()

	err := sqlite.UseMigrations(database)
	if err != nil {
		log.Fatal(err)
	}

	mwRepo := middleware.CreateNewRepo(database)
	handlersRepo := handlers.CreateNewRepo(database)

	handlers.SetNewHandlers(handlersRepo)
	middleware.SetNewMiddleware(mwRepo)

	srv := &http.Server{
		Addr:    config.PORT,
		Handler: routes.SetRoutes(),
	}

	log.Println("Starting application on port " + srv.Addr)
	if srv.ListenAndServe() != nil {
		log.Fatalf("%v - Internal Server Error", http.StatusInternalServerError)
	}
}
