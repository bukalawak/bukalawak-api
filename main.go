package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Run Bukalawak API")

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Pong")
	})

	http.HandleFunc("/posts", func(w http.ResponseWriter, r *http.Request) {
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

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

type Post struct {
	Tanya string `json:"tanya"`
	Jawab string `json:"jawab"`
}
