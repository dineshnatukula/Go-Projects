package main

import (
	"errors"
	"flag"
	"fmt"
	"net/http"
	"scheduler/internal/handler"
	"scheduler/internal/midleware"
	"scheduler/internal/repo"
	"scheduler/internal/service"

	"github.com/gorilla/mux"
	L "gitlab.gametechlabs.net/ppbet/backend-services/logger"
)

func main() {
	// Parsing flags
	flag.Parse()

	// Initialize the Logger

	// Initialize the database connection
	// dbConfig := models.DBConfig{
	// 	DbHost: *dbHost,
	// 	DbUser: *dbUser,
	// 	DbPort: *dbPort,
	// 	DbName: *dbName,
	// 	DbPass: *dbPass,
	// }
	// Initiliaze the Database.
	// TODO: uncomment the following line and above lines when deployed...
	// schedulerDB := InitializeDatabase(dbConfig)

	// Creating a Repo layer with the db details
	repo, err := repo.NewSchedulerRepo(nil)
	if err != nil {
		L.L.Fatal("No Jobs Found in DB to Schedule...")
		panic("No Jobs Found in DB to Schedule...")
	}

	// Creating a Service layer with the repo details
	service := service.NewSchedulerService(repo)

	// Creating the handler layer with the service details
	controller := handler.NewSchedulerHandler(service)

	// Initialize the mux router for handling http requests
	router := mux.NewRouter()

	// Use Middleware to log the request and response time durations.
	router.Use(midleware.LogRequest)

	// Register routes for the controller (PlayerHandler)
	controller.RegisterRoutes(router)

	// Setup HTTP server with the router
	srv := &http.Server{
		Addr:    *port,
		Handler: router,
	}

	// Start the HTTP server
	// L.L.Info("Starting server...")
	fmt.Println("Server Stating...")
	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		// L.L.Fatal("Server failed to start", L.Error(err))
	}
}
