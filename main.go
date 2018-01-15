package main

import (
	
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

// Init books as a slice book Struct
var books []Book



func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {

}

func createBook(w http.ResponseWriter, r *http.Request) {

}

func updateBook(w http.ResponseWriter, r *http.Request) {

}

func deleteBook(w http.ResponseWriter, r *http.Request) {

}

func main() {
	// Initialize Router
	router := mux.NewRouter()

	// Mock Data - @TODO - add db.
	books = append(books, Book{ID: "1", Isbn: "554566", Title: "Smallest Book Ever", Author: &Author{
		FirstName: "Jane", LastName: "Doe"}})

	books = append(books, Book{ID: "2", Isbn: "123456", Title: "Jazz Etudes for Tenor Sax", Author: &Author{
		FirstName: "Tyree", LastName: "Barron"}})

	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
