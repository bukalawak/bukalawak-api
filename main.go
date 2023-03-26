package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	fmt.Println("Run Bukalawak API")

	router := chi.NewRouter()

	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Pong")
	})
	router.Get("/posts", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		posts := [3]Post{
			{
				"Daerah Paling Ramah di Amerika?",
				"Michi-gan",
			},
			{
				"Makanan Favorit Batman?",
				"Batagor",
			},
			{
				"Apa beda nya Soto sama Cotto?",
				"Soto pakai daging sapi, kalau Cotto pakai daging Capii",
			},
		}
		json.NewEncoder(w).Encode(posts)
	})

	if err := http.ListenAndServe(":8000", router); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

type Post struct {
	Tanya string `json:"tanya"`
	Jawab string `json:"jawab"`
}
