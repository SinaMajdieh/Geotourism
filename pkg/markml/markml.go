package markml

import (
	"github.com/SinaMajdieh/Geotourism/pkg/domModel"
	"github.com/SinaMajdieh/Geotourism/pkg/file_pkg"
	"github.com/SinaMajdieh/Geotourism/pkg/log"
	"github.com/gomarkdown/markdown"
	"html/template"
	"os"
	"strings"
)

const (
	mdDir      = "website/blog/docs/md"
	linkPrefix = "blog/"
)

// Remedy Read markdown files
func remedy() []os.DirEntry {
	mds, err := os.ReadDir(mdDir)
	log.Log_err(err, "Reading files in "+mdDir)
	return mds
}

func Read_posts() []*domModel.BlogPost {
	mds := remedy()
	markmls := make([]*domModel.BlogPost, 0)
	for _, md := range mds {
		name := file_pkg.RemoveFileExt(md.Name())
		markml := &domModel.BlogPost{
			Title: name,
			Link:  linkPrefix + strings.ReplaceAll(name, " ", "-"),
			Html:  create_markml(md),
		}
		markmls = append(markmls, markml)
	}
	return markmls
}
func create_markml(md os.DirEntry) template.HTML {
	path := file_pkg.Make_path([]string{mdDir}, md.Name(), "")
	html := MdToHtml(path)
	return template.HTML(string(html))
}
func MdToHtml(path string) []byte {
	content, err := os.ReadFile(path)
	log.Log_err(err, "Reading the content of "+path)
	content = markdown.NormalizeNewlines(content)
	html := markdown.ToHTML(content, nil, nil)
	return html
}
