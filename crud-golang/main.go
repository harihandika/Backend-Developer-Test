package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Product holds your product attribute
type Data struct {
	ID              string `json:"id"`
	Nama            string `json:"nama"`
	TanggalLahir    string `json:"tanggal_lahir"`
	Kewarganegaraan string `json:"kewarganegaraan"`
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Selamat datang di home page")
}

func allDatas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(Datas)
}

func singleData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	id := params["id"]

	for _, data := range Datas {
		if data.ID == id {
			json.NewEncoder(w).Encode(data)
			return
		}
	}
}

func createData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var data Data

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	Datas = append(Datas, data)
	json.NewEncoder(w).Encode(data)
}

func updateData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	id := params["id"]

	var data Data

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for i, d := range Datas {
		if d.ID == id {
			Datas[i].Nama = data.Nama
			Datas[i].TanggalLahir = data.TanggalLahir
			Datas[i].Kewarganegaraan = data.Kewarganegaraan
			json.NewEncoder(w).Encode(Datas[i])
			return
		}
	}
}

func deleteData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	id := params["id"]

	for i, p := range Datas {
		if p.ID == id {
			Datas = append(Datas[:i], Datas[i+1:]...)
			json.NewEncoder(w).Encode(p)
			return
		}
	}
}

func handleRequest() {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/", home)
	r.HandleFunc("/datas", allDatas).Methods("GET")
	r.HandleFunc("/datas/{id}", singleData).Methods("GET")
	r.HandleFunc("/datas", createData).Methods("POST")
	r.HandleFunc("/datas/{id}", updateData).Methods("PUT")
	r.HandleFunc("/datas/{id}", deleteData).Methods("DELETE")

	fmt.Println("Application running")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func main() {
	Datas = []Data{
		Data{ID: "1", Nama: "Hendra", TanggalLahir: "20-10-2002", Kewarganegaraan: "Belanda"},
		Data{ID: "2", Nama: "Shara", TanggalLahir: "10-03-1998", Kewarganegaraan: "Indonesia"},
	}

	handleRequest()
}

// Datas is a global variable to hold collection of datas
var Datas []Data
