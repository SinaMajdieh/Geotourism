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
	blogPosts       []*domModel.BlogPost
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
	Pages["attraction_trimmed"] = template.Must(template.ParseFiles(
		pages_dir+"base.html",
		pages_dir+"attraction.html",
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
	//blog
	Pages["blog"] = template.Must(template.ParseFiles(
		pages_dir+"base.html",
		pages_dir+"navbar.html",
		pages_dir+"footer.html",
		pages_dir+"blog.html",
	))
	//blog post
	Pages["blog_post"] = template.Must(template.ParseFiles(
		pages_dir+"base.html",
		pages_dir+"navbar.html",
		pages_dir+"footer.html",
		pages_dir+"markml.html",
	))
	Pages["blog_post_trimmed"] = template.Must(template.ParseFiles(
		pages_dir+"base.html",
		pages_dir+"markml.html",
	))
	//404
	Pages["404"] = template.Must(template.ParseFiles(
		pages_dir+"base.html",
		pages_dir+"navbar.html",
		pages_dir+"404.html",
		pages_dir+"footer.html",
	))
	//comments
	Pages["comments"] = template.Must(template.ParseFiles(
		pages_dir+"comment_section.html",
		pages_dir+"comments.html",
		pages_dir+"commdown.html",
	))
	//comments
	Pages["comments"] = template.Must(template.ParseFiles(
		pages_dir+"base.html",
		pages_dir+"comment_section.html",
		pages_dir+"comments.html",
		//pages_dir+"commdown.html",
	))
	Pages["send_comment"] = template.Must(template.ParseFiles(
		pages_dir+"base.html",
		pages_dir+"comment_section.html",
		//pages_dir+"comments.html",
		pages_dir+"commdown.html",
	))
	//navbar
	Pages["navbar"] = template.Must(template.ParseFiles(
		pages_dir+"base.html",
		pages_dir+"navbar.html",
	))
	Pages["footer"] = template.Must(template.ParseFiles(
		pages_dir+"base.html",
		pages_dir+"footer.html",
	))
	//search results
	Pages["sat_result"] = template.Must(template.ParseFiles(
		pages_dir+"base.html",
		pages_dir+"navbar.html",
		pages_dir+"footer.html",
		pages_dir+"sat_result.html",
		pages_dir+"sat_result_block.html",
	))
}
