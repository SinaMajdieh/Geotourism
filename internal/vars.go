package internal

import (
	"html/template"

	"github.com/SinaMajdieh/Geotourism/pkg/domModel"
)

const (
	DocList   = "website/doc_list"
	pages_dir = "web/pages/"
)

var (
	phenomena       = []string{"Erosional", "Sedimentary_Geology", "Structural_Geology", "Magmatic_and_Metamorphic", "Geological_Landscape", "Engineering_Geology", "Cultural_Geology"}
	templates       = template.Must(template.ParseGlob("web/pages/*"))
	Pages           map[string]*template.Template
	attractions     *domModel.Attractions
	attractionsList *domModel.AttractionsList
	intros          *domModel.Articles
)

func DeclarePages() {
	Pages = make(map[string]*template.Template)
	Pages["home"] = template.Must(template.ParseFiles(
		pages_dir+"base.html",
		pages_dir+"navbar.html",
		pages_dir+"footer.html",
		pages_dir+"home.html",
	))
	Pages["article"] = template.Must(template.ParseFiles(
		pages_dir+"base.html",
		pages_dir+"navbar.html",
		pages_dir+"footer.html",
		pages_dir+"article.html",
	))
	Pages["attractions"] = template.Must(template.ParseFiles(
		pages_dir+"base.html",
		pages_dir+"navbar.html",
		pages_dir+"footer.html",
		pages_dir+"attractions_list.html",
		pages_dir+"attraction_item_grid.html",
	))
	Pages["attraction"] = template.Must(template.ParseFiles(
		pages_dir+"base.html",
		pages_dir+"navbar.html",
		pages_dir+"footer.html",
		pages_dir+"attraction.html",
		// TODO : reinvent the wheel (comment section)
		// backend : pkg/commenting
		//frontend : comment.htm, comment.js
		//disabled comment for now
		//"web/pages/comment.html",
	))
	Pages["comment"] = template.Must(template.ParseFiles(
		pages_dir+"base.html",
		pages_dir+"comment.html",
	))
	//for all kinds of crazy stuff!
	Pages["panel"] = template.Must(template.ParseFiles(
		pages_dir+"base.html",
		pages_dir+"navbar.html",
		//pages_dir+"footer.html",  doesn't stick to the bottom
		pages_dir+"panel.html",
	))
	//for editing test purposes
	Pages["editing"] = template.Must(template.ParseFiles(
		pages_dir+"base.html",
		pages_dir+"navbar.html",
		pages_dir+"footer.html",
		pages_dir+"edit.html",
	))
	Pages["edit_attraction"] = template.Must(template.ParseFiles(
		pages_dir+"base.html",
		pages_dir+"navbar.html",
		pages_dir+"footer.html",
		pages_dir+"edit_attraction.html",
	))
	//for adding purposes
	Pages["add_attraction"] = template.Must(template.ParseFiles(
		pages_dir+"base.html",
		pages_dir+"navbar.html",
		pages_dir+"footer.html",
		pages_dir+"add_attraction.html",
	))
}
