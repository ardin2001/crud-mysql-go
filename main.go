package main

import (
	"context"
	"crud_mysql/config"
	"crud_mysql/mahasiswa"
	"crud_mysql/utils"
	"fmt"
	"log"
	"net/http"
)

// GetMahasiswa
func GetMahasiswa(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		ctx, cancel := context.WithCancel(context.Background())

		defer cancel()

		mahasiswas, err := mahasiswa.GetAll(ctx)

		if err != nil {
			fmt.Println(err)
		}

		utils.ResponseJSON(w, mahasiswas, http.StatusOK)
		return
	} else if r.Method == "DELETE" {
		ctx, cancel := context.WithCancel(context.Background())

		defer cancel()

		id := r.URL.Query().Get("id")
		mahasiswas, err := mahasiswa.DeleteById(id, ctx)

		if err != nil {
			fmt.Println(err)
		}

		utils.ResponseJSON(w, mahasiswas, http.StatusOK)
		return
	} else if r.Method == "POST" {
		fmt.Println("post data")
		ctx, cancel := context.WithCancel(context.Background())

		defer cancel()

		nim := r.FormValue("nim")
		fullname := r.FormValue("name")
		semester := r.FormValue("semester")
		mahasiswas, err := mahasiswa.Create(nim, fullname, semester, ctx)

		if err != nil {
			fmt.Println(err)
		}

		utils.ResponseJSON(w, mahasiswas, http.StatusOK)
		return
	}

	http.Error(w, "Tidak di ijinkan", http.StatusNotFound)
}
func main() {

	db, e := config.MySQL()

	if e != nil {
		log.Fatal(e)
	}

	eb := db.Ping()
	if eb != nil {
		panic(eb.Error())
	}

	http.HandleFunc("/mahasiswa", GetMahasiswa)

	err := http.ListenAndServe(":7000", nil)
	fmt.Println("Server running in port : 7000")
	if err != nil {
		log.Fatal(err)
	}
}
