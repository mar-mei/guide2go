package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Server() {
	serverImagesPath := Config.Options.ImagesPath
	fs := http.FileServer(http.Dir(serverImagesPath))
	addr := Config.Options.Hostname

	log.Printf("Listening on: %s", addr)
	log.Printf("Using %s folder as image path", serverImagesPath)

	r := mux.NewRouter()

	if Config.Options.ProxyImages {
		r.HandleFunc("/images/{id}", proxyImages)
	} else if Config.Options.TVShowImages {
		r.PathPrefix("/images/").Handler(http.StripPrefix("/images/", fs))
	}

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
