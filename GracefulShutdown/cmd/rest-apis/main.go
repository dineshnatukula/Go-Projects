package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

type Book struct {
	ID            string `json:"id"`
	Title         string `json:"title"`
	ISBN          string `json:"isbn"`
	AuthorDetails Author `json:"authorDetails"`
}

type Author struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Response struct {
	Status           string `json:"status,omitempty"`
	ErrorCdoe        string `json:"errorCode,omitempty"`
	ErrorDescription string `json:"errorDescription,omitempty"`
}

func (res Response) String() string {
	return fmt.Sprintf("\"Status:\" %s\n\"ErrorCode:\" %s\n\"ErrorDescription:\" %s", res.Status, res.ErrorCdoe, res.ErrorDescription)
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Fetch book details from the db.. using limit / offset.
	// Can prepare a CSV and send as response as well... Export Option from FE.
	w.WriteHeader(http.StatusInternalServerError)
	if err := json.NewEncoder(w).Encode(books); err != nil {
		res := Response{
			Status:           "Failure",
			ErrorCdoe:        "Internal Server Error",
			ErrorDescription: "Error Ocurred due to marshaling the objects",
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(res)
	}
}

func getBookByID(w http.ResponseWriter, r *http.Request) {
	// Logging the request with the required details....
	log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)

	defer func(start time.Time) {
		log.Printf("%s %s", "Books sent as response", time.Since(start))
	}(time.Now())

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	book := findBookByID(params)

	// Book Not Found...
	if book == nil {
		res := Response{
			Status:           "Failure",
			ErrorCdoe:        "Not Found",
			ErrorDescription: "Book Not Found...",
		}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(res)
		return
	}

	// Book found and sends success response if the encoding is successful or ISE otherwise.
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(book); err != nil {
		res := Response{
			Status:           "Failure",
			ErrorCdoe:        "Internal Server Error",
			ErrorDescription: "Error Ocurred due to marshaling the objects",
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res)
	}
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var book Book

	// Get the request body and unmarhsal to book
	w.WriteHeader(http.StatusInternalServerError)
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		res := Response{
			Status:           "Failure",
			ErrorCdoe:        "Internal Server Error",
			ErrorDescription: "Error in creating book, Please try again..",
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	book.ID = strconv.Itoa(rand.Intn(100000))
	books = append(books, book)
	log.Printf("Book Created %s", book)
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(Response{Status: "Book Created Successfully..."}); err != nil {
		res := Response{
			Status:           "Failure",
			ErrorCdoe:        "Internal Server Error",
			ErrorDescription: "Error in creating book, Please try again..",
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res)
	}
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	var book *Book
	for idx, b := range books {
		if b.ID == params["id"] {
			book = &b
			books = append(books[:idx], books[idx+1:]...)
			break
		}
	}

	// Book Not Found...
	if book == nil {
		res := Response{
			Status:           "Failure",
			ErrorCdoe:        "Not Found",
			ErrorDescription: "Book Not Found...",
		}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(res)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(Response{Status: "Book Deleted Successfully..."}); err != nil {
		res := Response{
			Status:           "Failure",
			ErrorCdoe:        "Internal Server Error",
			ErrorDescription: "Book May be deleted... Please try to fetch with deleted BookID",
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res)
	}
}

func findBookByID(params map[string]string) (book *Book) {
	for _, b := range books {
		if b.ID == params["id"] {
			book = &b
			return
		}
	}
	return
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	book := findBookByID(params)
	// Book Not Found...
	if book == nil {
		res := Response{
			Status:           "Failure",
			ErrorCdoe:        "Not Found",
			ErrorDescription: "Book Not Found...",
		}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(res)
		return
	}
	id := book.ID
	json.NewDecoder(r.Body).Decode(&book)
	for idx, b := range books {
		if b.ID == params["id"] {
			books = append(books[:idx], books[idx+1:]...)
			var newBook Book
			json.NewDecoder(r.Body).Decode(&newBook)
			newBook.ID = params["id"]
			books = append(books, newBook)
			json.NewEncoder(w).Encode(newBook)
			return
		}
	}
	// ID is a primary key and won't be udpated...
	book.ID = id
	// Book found and updated successfully. Sends success response if the encoding is successful or ISE otherwise.
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(book); err != nil {
		res := Response{
			Status:           "Failure",
			ErrorCdoe:        "Internal Server Error",
			ErrorDescription: "Error Ocurred due to marshaling the objects",
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res)
	}
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Logging the request with the required details....
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)

		defer func(start time.Time) {
			log.Printf("%s %s", "Request and Resonse Duration", time.Since(start))
		}(time.Now())

		next.ServeHTTP(w, r)
	})
}

var books []Book

func main() {
	r := mux.NewRouter()
	r.Use(loggingMiddleware)

	books = append(books, Book{ID: "1", Title: "C Programming", ISBN: "123-asdf-23-sd-23", AuthorDetails: Author{ID: "1", FirstName: "Dennis", LastName: "Ritche"}})
	books = append(books, Book{ID: "2", Title: "Go Programming", ISBN: "124-asdf-23-sd-23", AuthorDetails: Author{ID: "1", FirstName: "Kevin", LastName: "Vyene"}})
	books = append(books, Book{ID: "3", Title: "Java Programming", ISBN: "125-asdf-23-sd-23", AuthorDetails: Author{ID: "1", FirstName: "Radu", LastName: "Alex"}})

	r.HandleFunc("/getBooks", getBooks).Methods("GET")
	r.HandleFunc("/getBook/{id}", getBookByID).Methods("GET")
	r.HandleFunc("/createBook", createBook).Methods("POST")
	r.HandleFunc("/deleteBook/{id}", deleteBook).Methods("DELETE")
	r.HandleFunc("/udpateBook/{id}", updateBook).Methods("PUT")

	log.Println("Server starting on port 8880")

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		fmt.Println("Starting server on :8080")
		// http.ListenAndServe(":8880", r)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("ListenAndServe error: %v\n", err)
		}
	}()

	// Set up channel on which to send signal notifications
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Block until a signal is received
	<-quit
	fmt.Println("Shutting down server...")

	// Create a deadline to wait for ongoing requests
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// The server is shutdown after the context is timedout.
	// The context is passed to the srv object.
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Printf("Server forced to shutdown: %v\n", err)
	}

	// Graceful Closing of DB Connections goes here...
	// Graceful Closing of Kafka Conenctions goes here...
	// Graceful Closing of Redis Connectiosn goes here...
	// Graceful Closing of any other conenctions...

	fmt.Println("Server exited gracefully")
}
