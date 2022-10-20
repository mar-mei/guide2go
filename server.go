package main

import (
	"log"
	"net/http"
	"os"
)

func Server() {
	port := os.Getenv("PORT")
	path := os.Getenv("IMAGES_PATH")
	fs := http.FileServer(http.Dir(os.Getenv("data/images")))
	log.Printf("Listening on: %s", port)
	log.Printf("using path: %s", path)
	addr := ":" + port
	log.Printf("using addr: %s", addr)
	err := http.ListenAndServe(addr, fs)
	if err != nil {
		log.Fatal(err)
	}
}
