package main

import (
	"github.com/SinaMajdieh/Geotourism/pkg/filePkg"
	"github.com/SinaMajdieh/Geotourism/pkg/tourson"
	"log"
	"net/http"
)
type HttpServer struct{
	IP   string `json:"ip"`
	Port string `json:"port"`
}
func (server HttpServer) SetupServer(){
	server.SetupAssets()
	server.SetupRoutes()
	server.ListenAndServe()
}
func (server HttpServer) SetupRoutes() {
	http.HandleFunc("/home", makeHandler(indexHandler))
	http.HandleFunc("/articles/", makeHandler(articleHandler))
	http.HandleFunc("/attraction/", makeHandler(attractionHandler))
	http.HandleFunc("/attractions", makeHandler(attractionsHandler))
	http.HandleFunc("/intro", makeHandler(introHandler))
}
func (server HttpServer)ListenAndServe(){
	log.Println("Listening on :80...")
	err := http.ListenAndServe(server.IP + server.Port, nil)
	if err != nil {
		log.Fatal(err)
	}

}
func (server HttpServer)SetupAssets(){
	css := http.FileServer(http.Dir("./web/stylesheets"))
	http.Handle("/css/", http.StripPrefix("/css/" , css))
	js := http.FileServer(http.Dir("./web/javascripts"))
	http.Handle("/js/", http.StripPrefix("/js/" , js))
	img := http.FileServer(http.Dir("./web/images"))
	http.Handle("/img/", http.StripPrefix("/img/" , img))
	fonts := http.FileServer(http.Dir("./web/fonts"))
	http.Handle("/fonts/", http.StripPrefix("/fonts/" , fonts))
	docImg := http.FileServer(http.Dir("./Articles/Img"))
	http.Handle("/docimg/", http.StripPrefix("/docimg/" , docImg))
}
func (server *HttpServer)Initialize(configFile string){
	err := filePkg.ReadJson(configFile , server)
	if nil != err {
		log.Fatal(err)
	}
}

func LoadData() {
	attractionsList , attractions = tourson.LoadAttractionsListFiles(phenomena)
	intros = tourson.LoadIntroListFiles(DocList)
}
