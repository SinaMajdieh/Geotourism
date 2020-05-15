package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var (
	templates = template.Must(template.ParseGlob("web/Pages/*"))
	home      = template.Must(template.ParseFiles(
		"web/Pages/base.html",
		"web/Pages/navbar.html",
		"web/Pages/footer.html",
		"web/Pages/home.html",
	))
	articlePage = template.Must(template.ParseFiles(
		"web/Pages/base.html",
		"web/Pages/navbar.html",
		"web/Pages/footer.html",
		"web/Pages/article.html",
	))
	attractionsPage = template.Must(template.ParseFiles(
		"web/Pages/base.html",
		"web/Pages/navbar.html",
		"web/Pages/footer.html",
		"web/Pages/attractionsList.html",
		"web/Pages/attractionItemGrid.html",
	))
	attractionPage = template.Must(template.ParseFiles(
		"web/Pages/base.html",
		"web/Pages/navbar.html",
		"web/Pages/footer.html",
		"web/Pages/attraction.html",
	))
	attractions     *Attractions
	attractionsList *AttractionsList
	intros          *Articles
)

func makeHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r)
	}

}
func main() {
	LoadDatas()
	var server Server
	server = &HttpServer{}
	server.Initialize("ServerConfig.json")
	server.ListenAndServe()

}
func renderTemplate(w http.ResponseWriter, tmpl string, article *Article) {

	err := templates.ExecuteTemplate(w, tmpl+".html", article)

	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)

	}

}
func indexHandler(w http.ResponseWriter, r *http.Request) {

	err := home.Execute(w, nil)
	if nil != err {
		fmt.Println(err)
	}

}
func articleHandler(w http.ResponseWriter, r *http.Request) {
	articleName := r.URL.Path[len("/articles/"):]
	var article Article
	path := MakePath([]string{GlobalArticlesDirectory, articleName}, articleName, "json")
	err := ReadJson(path, &article)
	if nil != err {
		fmt.Println(err)
	}
	title := article.Title
	article.Title = ""
	articles := Articles{
		Title:    title,
		Articles: []Article{article},
	}
	articlePage.Execute(w, articles)

}
func attractionHandler(w http.ResponseWriter, r *http.Request) {
	link := r.URL.Path
	link = link[1:]
	attraction := FindAttraction(link)
	if nil == attraction {
		fmt.Println("Not Found")
	} else {
		attractionPage.Execute(w, attraction)
		fmt.Println(attraction.MapImage)
	}

}
func introHandler(w http.ResponseWriter, r *http.Request) {
	articlePage.Execute(w, intros)
}
func attractionsHandler(w http.ResponseWriter, r *http.Request) {
	_ = attractionsPage.Execute(w, attractionsList)

}
func downloadHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w , r , "Articles/Attractions/Docs/Chah Dashi Marn.json")

}
