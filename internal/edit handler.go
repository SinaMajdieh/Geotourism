package internal

import (
	"fmt"
	"net/http"

	"github.com/TwiN/go-color"
)

func edit_list_handler(w http.ResponseWriter, r *http.Request) {
	err := Pages["editing"].Execute(w, attractionsList)
	if nil != err {
		fmt.Println(err)
	}
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
