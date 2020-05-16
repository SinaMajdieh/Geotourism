package internal

import (
	"fmt"
	"github.com/SinaMajdieh/Geotourism/pkg/domModel"
	"github.com/SinaMajdieh/Geotourism/pkg/filePkg"
	"net/http"
)

var (

)

func makeHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r)
	}

}
func renderTemplate(w http.ResponseWriter, tmpl string, article *domModel.Article) {

	err := templates.ExecuteTemplate(w, tmpl+".html", article)

	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)

	}

}
func indexHandler(w http.ResponseWriter, r *http.Request) {

	err := Pages["home"].Execute(w, nil)
	if nil != err {
		fmt.Println(err)
	}

}
func articleHandler(w http.ResponseWriter, r *http.Request) {
	articleName := r.URL.Path[len("/articles/"):]
	var article domModel.Article
	path := filePkg.MakePath([]string{filePkg.GlobalArticlesDirectory, articleName}, articleName, "json")
	err := filePkg.ReadJson(path, &article)
	if nil != err {
		fmt.Println(err)
	}
	title := article.Title
	article.Title = ""
	articles := domModel.Articles{
		Title:    title,
		Articles: []domModel.Article{article},
	}
	Pages["article"].Execute(w, articles)

}
func attractionHandler(w http.ResponseWriter, r *http.Request) {
	link := r.URL.Path
	link = link[1:]
	attraction := FindAttraction(link)
	if nil == attraction {
		fmt.Println("Not Found")
	} else {
		Pages["attraction"].Execute(w, attraction)
	}

}
func introHandler(w http.ResponseWriter, r *http.Request) {
	Pages["article"].Execute(w, intros)
}
func attractionsHandler(w http.ResponseWriter, r *http.Request) {
	_ = Pages["attractions"].Execute(w, attractionsList)

}