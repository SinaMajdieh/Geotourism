package main

import (
	"log"
	"net/http"
)
type HttpServer struct{
	IP   string `json:"ip"`
	Port string `json:"port"`
}
func (server HttpServer) ListenAndServe() {
	fs := http.FileServer(http.Dir("."))
	http.Handle("/", fs)
	http.HandleFunc("/home", makeHandler(indexHandler))
	http.HandleFunc("/articles/", makeHandler(articleHandler))
	http.HandleFunc("/attraction/", makeHandler(attractionHandler))
	http.HandleFunc("/attractions", makeHandler(attractionsHandler))
	http.HandleFunc("/intro", makeHandler(introHandler))
	http.HandleFunc("/download", makeHandler(downloadHandler))
	log.Println("Listening on :80...")
	err := http.ListenAndServe(server.IP + server.Port, nil)
	if err != nil {
		log.Fatal(err)
	}

}
func (server *HttpServer)Initialize(configFile string){
	err := ReadJson(configFile , server)
	if nil != err {
		log.Fatal(err)
	}
}

func LoadDatas() {
	attractionsList = LoadAttractionsListFiles()
	intros = LoadIntroListFiles()
}
