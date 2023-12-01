package internal

import (
	"Geotourism/pkg/domModel"
	"Geotourism/pkg/log"
	"Geotourism/pkg/tourson"
	"github.com/TwiN/go-color"
	"github.com/gorilla/schema"
	"net/http"
	"strconv"
	"strings"
)

func edit_list_handler(w http.ResponseWriter, r *http.Request) {
	err := Pages["editing"].Execute(w, attractionsList)
	log.Log_err(err, "")
}

func edit_attraction_handler(w http.ResponseWriter, r *http.Request) {
	prefix := "edit/attraction/"
	link := r.URL.Path
	link = link[len(prefix):]
	attraction := FindAttraction("attraction" + link)
	if nil == attraction {
		println(color.Ize(color.Red, "couldn't find "+link))
	} else {
		Pages["edit_attraction"].Execute(w, attraction)
	}
}

func editing_result_handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.Redirect(w, r, "/edit/attractions", http.StatusSeeOther)
	} else {
		att := set_edited_attraction_attributes(r)
		err := tourson.Save_attraction_file(&att)
		log.Log_err(err, att.Title+" was edited successfully!")
		/*
			TODO make it work
			Update_attractions_list(&att)
		*/
		Load_attractions()
		http.Redirect(w, r, "/edit/attractions", http.StatusSeeOther)
	}
}

func set_edited_attraction_attributes(r *http.Request) domModel.Attraction {
	err := r.ParseForm()
	log.Log_err(err, "")
	var att domModel.Attraction

	att = domModel.Attraction{}
	decoder := schema.NewDecoder()
	_ = decoder.Decode(&att, r.PostForm)

	att.Content = strings.Split(att.Content[0], "\r\n")
	att.Seasons = make_seasons_arr(att.SeasonsToVisit)
	att.Gallery = set_pics(r)

	return att
}
func make_seasons_arr(seasons string) []string {
	seasons = strings.ReplaceAll(seasons, " ", "")
	seasons_arr := strings.Split(seasons, ",")
	return seasons_arr
}
func set_pics(r *http.Request) []domModel.Picture {
	pics := make([]domModel.Picture, 0)
	for i := 0; r.PostForm.Has("src" + strconv.Itoa(i)); i++ {
		pic := domModel.Picture{
			Src:     r.Form["src"+strconv.Itoa(i)][0],
			Caption: r.Form["cap"+strconv.Itoa(i)][0],
		}
		pics = append(pics, pic)
	}
	return pics
}
