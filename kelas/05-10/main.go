package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// template diluar handler
// setiap halaman punya set template sendiri, karena masing-masing
// mendefinisikan template "content" yang berbeda
var tplHome = template.Must(template.New("layout").ParseFiles(
	"template/layout.html",
	"template/home.html",
))

var tplProduk = template.Must(template.New("layout").ParseFiles(
	"template/layout.html",
	"template/produk.html",
))

var tplProdukForm = template.Must(template.New("layout").ParseFiles(
	"template/layout.html",
	"template/produk_form.html",
))

type User struct {
	Name  string
	Email string
	Age   int
}

var DataUser []User = []User{
	User{Name: "Anggyi Trisnawan", Email: "anggyi@gmail.com", Age: 32},
	User{Name: "Anggi Putra", Email: "anggi@gmail.com", Age: 34},
}

type Produk struct {
	ID    int
	Nama  string
	Harga string
}

// penyimpanan di memori program
var DataProduk []Produk
var nextID = 1

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", homeHandler)
	mux.HandleFunc("GET /produk", produkHandler)
	mux.HandleFunc("GET /produk/form", produkFormHandler)
	mux.HandleFunc("POST /produk/simpan", produkSimpanHandler)
	http.ListenAndServe(":8080", mux)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]any{
		"Judul": "Halaman Home",
		"Hewan": []string{"Kambing", "Kuda", "Kerbau"},
		"Users": DataUser,
	}
	if err := tplHome.ExecuteTemplate(w, "layout", data); err != nil {
		fmt.Println("Ada Error", err)
	}
}


func produkHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]any{
		"Judul":  "Daftar Produk",
		"Produk": DataProduk,
	}
	if err := tplProduk.ExecuteTemplate(w, "layout", data); err != nil {
		fmt.Println("Ada Error", err)
	}
}

func produkFormHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]any{
		"Judul": "Tambah Produk",
	}
	if err := tplProdukForm.ExecuteTemplate(w, "layout", data); err != nil {
		fmt.Println("Ada Error", err)
	}
}

func produkSimpanHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Form tidak valid", http.StatusBadRequest)
		return
	}

	nama := r.FormValue("nama")
	harga := r.FormValue("harga")

	if nama == "" || harga == "" {
		data := map[string]any{
			"Judul": "Tambah Produk",
			"Error": "Nama dan harga wajib diisi",
			"Nama":  nama,
			"Harga": harga,
		}
		w.WriteHeader(http.StatusBadRequest)
		tplProdukForm.ExecuteTemplate(w, "layout", data)
		return
	}

	DataProduk = append(DataProduk, Produk{ID: nextID, Nama: nama, Harga: harga})
	nextID++

	http.Redirect(w, r, "/produk", http.StatusSeeOther)
}
