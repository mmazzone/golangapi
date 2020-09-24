package main

import (
	"encoding/json"
	"log"
	"net/http"
	"math/rand"
	"strconv"
	"git.com/gorilla/mux"
)

// Book struct (model)
type Book struct {
	ID		string `json:"id"`
	Isbn	string `json:"isbn"`
	Title	string `json:"title"`
	Author	*Author `json:"author"`
}

// Author struct
type Author struct {
	Firstname	string `json:"firstname"`
	Lastname	string `json:"lastname"`
}

// INIT BOOKS VAR AS A SLICE BOOK STRUCT
var books []Book

// GET ALL BOOKS
funct getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}
// GET SINGLE BOOK
funct getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // GET PARAMS
	// LOOP THROUGH BOOKS AND FIND WITH ID
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

// CREATE A NEW BOOK
funct createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000)) // MOCK ID
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}
// UPDATE BOOK
funct updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.Id == params["id"] {
			books = append(books[:index], books[index+1]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(books)
}
// DELETE BOOK
funct deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.Id == params["id"] {
			books = append(books[:index], books[index+1]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}


func main() {
	// Init router
	r := mux.NewRouter()

	// MOCK DATA - @TODO - IMPLEMENT DB
	books = append(books, Book{ID: "1", Isbn:"345345", Title: " Book One", Author: &Author
	{Firtsname: "John", Lastname: "Doe"}})
	books = append(books, Book{ID: "2", Isbn:"678678", Title: " Book One", Author: &Author
	{Firtsname: "John", Lastname: "Doe"}})
	

	// Route Handlers / end points
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBooks).Methods("GET")
	r.HandleFunc("/api/books", createBooks).Methods("POST")
	r.HandleFunc("/api/books{id}", updateBooks).Methods("PUT")
	r.HandleFunc("/api/books{id}", deleteBooks).Methods("DELETE")
	
	log.Fatal(http.ListenAndServe(":8000", r))
}

