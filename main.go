package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func main() {
	fmt.Println("Run Bukalawak API")

	router := chi.NewRouter()

	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Pong")
	})

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

	router.Get("/posts", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(posts)
	})

	router.Get("/post/{postId}", func(w http.ResponseWriter, r *http.Request) {
		postId, err := strconv.Atoi(chi.URLParam(r, "postId"))
		if err != nil {
			w.WriteHeader(404)
			io.WriteString(w, "Not Found")
			return
		}

		// sanitize
		postId = postId - 1
		if postId < 0 || postId >= len(posts) {
			w.WriteHeader(404)
			io.WriteString(w, "Not Found")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(posts[postId])
	})

	if err := http.ListenAndServe(":8000", router); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

type Post struct {
	Tanya string `json:"tanya"`
	Jawab string `json:"jawab"`
}
