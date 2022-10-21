package main

import (
	"log"
	"net/http"
	"os"
)

func Server() {
	port := os.Getenv("PORT")
	serverImagesPath := Config.Options.ImagesPath
	fs := http.FileServer(http.Dir(serverImagesPath))
	log.Printf("Listening on: %s", port)
	log.Printf("Using %s folder as image path", serverImagesPath)
	addr := ":" + port
	err := http.ListenAndServe(addr, fs)
	if err != nil {
		log.Fatal(err)
	}
}
