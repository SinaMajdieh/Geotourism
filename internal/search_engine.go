package internal

import (
	"net/http"
)

func search(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.Redirect(w, r, "/home", http.StatusFound)
	} else {
		err := r.ParseForm()
		if nil != err{
			Pages["404"].Execute(w, nil)
		}
		index := r.Form["index"]
		searchResult := Sattaction(index[0])

		if len(searchResult) == 0 {
			Pages["404"].Execute(w, nil)
		} else {
			Pages["sat_result"].Execute(w, searchResult)
		}
	}
}
