package main

import "flag"

var (
	dbName = flag.String("db.name", "scheduler", "Database name for Scheduler service talks to")
	dbHost = flag.String("db.host", "localhost", "Address that db is hosted on")
	dbPort = flag.String("db.port", "5432", "Port number that the db is running on")
	dbUser = flag.String("db.user", "postgres", "Scheduler database username")
	dbPass = flag.String("db.pass", "ppbet123", "password of the scheduler db")
	port   = flag.String("port", ":8880", "Service Port")
	// internalAPIsPort = flag.String("internalapis.port", "8880", "Internal APIs listening port")
)
