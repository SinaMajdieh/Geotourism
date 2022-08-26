package internal

import (
	"github.com/SinaMajdieh/Geotourism/pkg/log"
	"github.com/SinaMajdieh/Geotourism/pkg/tourson"
	"github.com/TwiN/go-color"
	"net/http"
)

func remove_handler(w http.ResponseWriter, r *http.Request) {
	err := Pages["editing"].Execute(w, attractionsList)
	log.Log_err(err, "")

}
func remove_result_handler(w http.ResponseWriter, r *http.Request) {
	prefix := "remove/attraction/"
	link := r.URL.Path
	link = link[len(prefix):]
	attraction := FindAttraction("attraction" + link)
	if nil == attraction {
		println(color.Ize(color.Red, "couldn't find "+link))
	} else {
		err := tourson.Remove_attraction_file(attraction.Title)
		log.Log_err(err, attraction.Title+" removed successfully.")
		Load_attractions()
		http.Redirect(w, r, "/remove/attraction", http.StatusSeeOther)
	}
}
