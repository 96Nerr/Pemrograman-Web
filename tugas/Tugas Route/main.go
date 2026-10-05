package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	
	mux.HandleFunc("GET /{$}", HomeHandler)
	mux.HandleFunc("GET /profile", ProfileHandler)
	mux.HandleFunc("GET /projects", ProjectsHandler)
	mux.HandleFunc("GET /portofolio", PortofolioHandler)
	mux.HandleFunc("GET /articles/{id}", ArticlesHandler)
	mux.HandleFunc("POST /articles/{id}", ArticlesHandler)

	http.ListenAndServe(":8000", mux)

	fmt.Println("Server berjalan di http://localhost:8000")

	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		fmt.Println("Server error:", err)
	}

}


func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the Home Page")
}

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "About Me")
}


func PortofolioHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "=== PORTOFOLIO ===")
	fmt.Fprintln(w, "Nama: Naufal")
	fmt.Fprintln(w, "Prodi: Teknik Informatika")
	fmt.Fprintln(w, "Project: Website, OpenGL, dan Machine Learning")
}

func ProjectsHandler(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")

	if page == "" {
		page = "1"
	}

	fmt.Fprintln(w, "Halaman:", page)
}

func ArticlesHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	if r.Method == "POST" {
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintln(w, "Article berhasil dibuat dengan ID:", id)
		return
	}

	fmt.Fprintln(w, "Article ID:", id)
}

