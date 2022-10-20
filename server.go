package main

import (
	"log"
	"net/http"
	"os"
)

func Server() {
	port := os.Getenv("PORT")
	fs := http.FileServer(http.Dir("/data/images"))
	log.Printf("Listening on: %s", port)
	addr := ":" + port
	err := http.ListenAndServe(addr, fs)
	if err != nil {
		log.Fatal(err)
	}
}
