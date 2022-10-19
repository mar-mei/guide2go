package main

import (
	"log"
	"net/http"
	"os"
)

func Server() {
	port := os.Getenv("PORT")
	fs := http.FileServer(http.Dir(os.Getenv("IMAGES_PATH")))
	http.Handle("/", fs)

	log.Printf("Listening on :%s...", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
