package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"os"

	"github.com/gorilla/mux"
)

func Server() {
	log.SetOutput(os.Stdout)
	port := strings.Split(Config.Options.Hostname, ":")
	var addr string
	serverImagesPath := Config.Options.ImagesPath
	fs := http.FileServer(http.Dir(serverImagesPath))
	if len(port) == 2 {
		addr = ":" + port[1]
	} else {
		log.Println("No port found, using port 8080")
		addr = ":8080"
	}

	log.Printf("Listening on: %s", addr)
	log.Printf("Using %s folder as image path", serverImagesPath)

	r := mux.NewRouter()

	if Config.Options.ProxyImages {
		r.HandleFunc("/images/{id}", proxyImages)
	} else if Config.Options.TVShowImages {
		r.PathPrefix("/images/").Handler(http.StripPrefix("/images/", fs))
	}
	r.HandleFunc("/run", run)

	err := http.ListenAndServe(addr, r)
	if err != nil {
		log.Fatal(err)
	}
}

func proxyImages(w http.ResponseWriter, r *http.Request) {
	image := mux.Vars(r)
	url := "https://json.schedulesdirect.org/20141201/image/" + image["id"] + "?token=" + Token
	a, _ := http.NewRequest("GET", url, nil)
	http.Redirect(w, a, url, http.StatusSeeOther)
	log.Println("requested image: " + r.RequestURI)
}

func run(w http.ResponseWriter, r *http.Request) {
	var sd SD
	go sd.Update(Config2)
	fmt.Fprint(w, "Grabbing EPG")
}
