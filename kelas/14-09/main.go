package main

import (

	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	
	mux.HandleFunc("GET /", HomeHandler)
	mux.HandleFunc("GET /about", AboutHandler)
	mux.HandleFunc("GET /contact", ContactHandler)
	mux.HandleFunc("GET /products/{id...}", ProductsHandler)
	http.ListenAndServe(":8000", mux)

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "About Us")
}

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	fmt.Fprintf(w, "Product ID: %s", id)
}

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Contact Us")
}