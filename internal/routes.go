package internal

import (
	"github.com/SinaMajdieh/Geotourism/pkg/commenting"
	"github.com/SinaMajdieh/Geotourism/pkg/tourson"
	"log"
	"net/http"
)

type HttpServer struct {
	IP   string `json:"ip"`
	Port string `json:"port"`
}

func (server HttpServer) SetupServer() {
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
	//testing route for commenting
	http.HandleFunc("/comment", makeHandler(comment_handler))
	http.HandleFunc("/comment-echo", makeHandler(commenting.Comment_handler))
	http.HandleFunc("/comment-load", makeHandler(commenting.Comment_load_handler))
}
func (server HttpServer) ListenAndServe() {
	log.Println("Listening on : " + server.IP + server.Port)
	err := http.ListenAndServe(server.IP+server.Port, nil)
	if err != nil {
		log.Fatal(err)
	}

}
func (server HttpServer) SetupAssets() {
	css := http.FileServer(http.Dir("./web/stylesheets"))
	http.Handle("/css/", http.StripPrefix("/css/", css))
	js := http.FileServer(http.Dir("./web/javascripts"))
	http.Handle("/js/", http.StripPrefix("/js/", js))
	img := http.FileServer(http.Dir("./web/images"))
	http.Handle("/img/", http.StripPrefix("/img/", img))
	gif := http.FileServer(http.Dir("./web/gifs"))
	http.Handle("/gif/", http.StripPrefix("/gif/", gif))
	fonts := http.FileServer(http.Dir("./web/fonts"))
	http.Handle("/fonts/", http.StripPrefix("/fonts/", fonts))
	docImg := http.FileServer(http.Dir("./website/articles/img"))
	http.Handle("/docimg/", http.StripPrefix("/docimg/", docImg))
}
func NewHttpServer(ip string, port string) *HttpServer {
	return &HttpServer{}
}

func LoadData() {
	attractionsList, attractions = tourson.LoadAttractionsListFiles(phenomena)
	intros = tourson.LoadIntroListFiles(DocList)
}
