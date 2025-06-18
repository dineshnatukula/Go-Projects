package main

import (
	"fmt"
	"log"
	"scheduler/internal/models"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	// L "gitlab.gametechlabs.net/ppbet/backend-services/logger"
)

func InitLogging() {
	// L.Init()
	// defer L.L.Sync()
	// L.L.Info("Logger Initialization done...")
}

func InitMetrics() {
	// Initialisation of the Mux...
	// L.L.Info("MUX Initilization Started.")
	// M.M.InitAll("playerriskmanager", ":"+*internalAPIsPort, nil, nil, nil, nil, []string{"apirequests_count", "apirequests_durationMS", "", "version", "method", "status", "functionName"})
	// M.M.Mux.Handle("/loglevel", L.L.Config.Level)
	// M.M.Mux.HandleFunc("/log", L.L.GetLog)
	// M.M.Mux.Handle("/health", health.Handler())
	// pprof api

	// M.M.Mux.HandleFunc("/debug/pprof/", pprof.Index)
	// M.M.Mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	// M.M.Mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	// M.M.Mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	// M.M.Mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	// L.L.Info("Metrics initialization done...")
}

func InitializeDatabase(dbConfig models.DBConfig) (db *gorm.DB) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", dbConfig.DbHost, dbConfig.DbUser, dbConfig.DbPass, dbConfig.DbName, dbConfig.DbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		// L.Fatal("Failed to connect to database: %v", L.Error(err))
		panic(err)
	}

	// / Get the generic database object sql.DB to use its functions
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get sql.DB from gorm DB: %v", err)
	}

	// Configure connection pool
	sqlDB.SetMaxOpenConns(100)                // Maximum number of open connections to the database
	sqlDB.SetMaxIdleConns(100)                // Maximum number of idle connections in the pool
	sqlDB.SetConnMaxLifetime(5 * time.Minute) // Maximum amount of time a connection may be reused
	return
}
