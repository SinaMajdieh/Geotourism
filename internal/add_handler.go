package internal

import (
	"Geotourism/pkg/domModel"
	"Geotourism/pkg/log"
	"Geotourism/pkg/tourson"
	"net/http"
)

func add_attraction_handler(w http.ResponseWriter, r *http.Request) {
	Pages["add_attraction"].Execute(w, nil)
}
func add_result_handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.Redirect(w, r, "/add/attraction", http.StatusSeeOther)
	} else {
		att := set_added_attraction_attributes(r)
		err := tourson.Save_attraction_file(&att)
		log.Log_err(err, att.Title+" was added successfully")
		/*
			TODO make it work
			Update_attractions_list(&att)
		*/
		Load_attractions()
		http.Redirect(w, r, "/add/attraction", http.StatusSeeOther)
	}
}
func set_added_attraction_attributes(r *http.Request) domModel.Attraction {
	return set_edited_attraction_attributes(r)
}
