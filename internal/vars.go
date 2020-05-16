package internal

import (
	"github.com/SinaMajdieh/Geotourism/pkg/domModel"
	"html/template"
)

const (
	DocList = "website/doc_list"
)

var (
	phenomena = []string{"Erosional" , "Sedimentary" , "Structural"}
	templates = template.Must(template.ParseGlob("web/pages/*"))
	Pages map[string]*template.Template
	attractions *domModel.Attractions
	attractionsList *domModel.AttractionsList
	intros *domModel.Articles
)

func DeclarePages(){
	Pages = make(map[string]*template.Template)
	Pages["home"] = template.Must(template.ParseFiles(
		"web/pages/base.html",
		"web/pages/navbar.html",
		"web/pages/footer.html",
		"web/pages/home.html",
	))
	Pages["article"] = template.Must(template.ParseFiles(
		"web/pages/base.html",
		"web/pages/navbar.html",
		"web/pages/footer.html",
		"web/pages/article.html",
	))
	Pages["attractions"] = template.Must(template.ParseFiles(
		"web/pages/base.html",
		"web/pages/navbar.html",
		"web/pages/footer.html",
		"web/pages/attractionsList.html",
		"web/pages/attractionItemGrid.html",
	))
	Pages["attraction"] = template.Must(template.ParseFiles(
		"web/pages/base.html",
		"web/pages/navbar.html",
		"web/pages/footer.html",
		"web/pages/attraction.html",
	))
}