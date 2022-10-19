package main

import (
	"log"
	"net/http"
)

func Server(port string) {
	fs := http.FileServer(http.Dir("/Users/jesusdavid/Documents/images"))
	http.Handle("/", fs)

	log.Printf("Listening on :%s...", port)
	err := http.ListenAndServe(":"+ port, nil)
	if err != nil {
		log.Fatal(err)
	}
}