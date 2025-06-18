package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

type User struct {
	ID        string `json:"id,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"secondName,omitempty"`
}

// GetUsers returns a list of users
// @Summary Get a list of users
// @Description Get all users
// @Produce json
// @Success 200 {array} User
// @Failure 500 Internal Server Error
// @Router /users [GET]
func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := []User{
		{ID: "1", FirstName: "Dinesh", LastName: "Natukula"},
		{ID: "2", FirstName: "Jaya Prakash", LastName: "Aluri"},
	}

	// May be fetch all the data based on limit and offset or based on a date range from db.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

// GetUserByID returns details of a specific user by ID
// @Summary Get user by ID
// @Description Get details of a specific user
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} User
// @Router /users/{id} [get]
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	user := User{ID: userID, FirstName: fmt.Sprintf("User %s", userID)}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

// @title APISwagger
// @version 1.0
// @description This is a sample REST API to demonstrate in generating the Swagger documentation.
// @host localhost:8080
// @BasePath /api/v1
func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/v1/getusers", GetUsers).Methods("GET")
	router.HandleFunc("/api/v1/getusers/{id}", GetUserByID).Methods("GET")
	// router.PathPrefix("/swagger").Handler(httpSwagger.Handler(httpSwagger.URL("/swagger/doc.json")))
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// http.Handle("/", router)
	http.ListenAndServe(":8080", router)
}
