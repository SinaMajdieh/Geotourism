package internal

import (
	"fmt"
	"Geotourism/pkg/commdown"
	"Geotourism/pkg/domModel"
	"net/http"
)

func blog_handler(w http.ResponseWriter, r *http.Request) {
	Pages["blog"].Execute(w, blogPosts)
}

// original handler with no comment loading
/*func blog_post_handler(w http.ResponseWriter, r *http.Request) {
	link := r.URL.Path[1:]
	post := find_post(link)
	if nil != post {
		Pages["blog_post"].Execute(w, post)
	} else {
		Pages["404"].Execute(w, nil)
	}
}*/

func blog_post_handler(w http.ResponseWriter, r *http.Request) {
	link := r.URL.Path[1:]
	fmt.Println(link)
	post := find_post(link)

	Pages["navbar"].Execute(w, nil)

	if nil != post {
		Pages["blog_post_trimmed"].Execute(w, post)
	} else {
		Pages["404"].Execute(w, nil)
	}

	comments := commdown.GetIt(link)
	if nil != comments {
		Pages["comments"].Execute(w, comments)
	} else {
		fmt.Println("Not found comments")
	}

	Pages["send_comment"].Execute(w, link)

	Pages["footer"].Execute(w, nil)
}

func find_post(link string) *domModel.BlogPost {
	for _, v := range blogPosts {
		if v.Link == link {
			return v
		}
	}
	return nil
}
