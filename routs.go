package main

import (
	"log"
	"net/http"
)

func SetupRouts() {
	fs := http.FileServer(http.Dir("."))
	http.Handle("/", fs)
	http.HandleFunc("/home", makeHandler(indexHandler))
	http.HandleFunc("/articles/", makeHandler(articleHandler))
	http.HandleFunc("/attraction/", makeHandler(attractionHandler))
	http.HandleFunc("/attractions", makeHandler(attractionsHandler))
	http.HandleFunc("/intro", makeHandler(introHandler))
	log.Println("Listening on :80...")
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func LoadDatas() {
	attractions = LoadAttarctionsListFiles()
	intros = LoadIntroListFiles()
}
