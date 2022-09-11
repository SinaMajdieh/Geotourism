package internal

import (
	"fmt"
	"github.com/SinaMajdieh/Geotourism/pkg/commdown"
	"github.com/SinaMajdieh/Geotourism/pkg/domModel"
	"github.com/SinaMajdieh/Geotourism/pkg/log"
	"github.com/gorilla/schema"
	"net/http"
)

func commentApiHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
		fmt.Println("huh")
	} else {
		preficx := "/commetn-api/"
		url := r.URL.Path[len(preficx):]
		fmt.Println(url)
		err := r.ParseForm()
		log.Log_err(err, "")

		newComment := domModel.Comment{}
		decoder := schema.NewDecoder()
		err = decoder.Decode(&newComment, r.PostForm)

		newComment.Link = url

		commdown.SaveComment(&newComment)
		//err = commlight.InsertComment(&newComment)
		log.Log_err(err, "comment inserted")

		http.Redirect(w, r, "/"+url, http.StatusSeeOther)
	}
}
