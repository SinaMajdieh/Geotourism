package commdown

import (
	"fmt"
	"Geotourism/pkg/domModel"
	"Geotourism/pkg/file_pkg"
	"Geotourism/pkg/log"
	"Geotourism/pkg/markml"
	"html/template"
	"os"
)

var (
	DataDownPath = "website/comments"
)

const (
	EndLine     = "  \n"
	Quote       = "> "
	EndSequence = "\n\n"
)

func SetDataDownPath(path string) {
	DataDownPath = path
}
func SaveComment(comment *domModel.Comment) {
	path := file_pkg.Make_path([]string{DataDownPath}, comment.Link, ".md")
	//path := DataDownPath + comment.Link + ".md"
	fmt.Println(path)
	if !file_pkg.FileExists(path) {
		createCommdownFile(path)
	}
	old := readIt(path)
	new := mdIt(comment) + old
	writeIt(path, new)

}
func createCommdownFile(path string) {
	file, err := os.Create(path)
	log.Log_err(err, "")
	file.Close()
}
func readIt(path string) string {
	content, err := os.ReadFile(path)
	log.Log_err(err, "")
	return string(content)
}
func writeIt(path, content string) {
	err := os.WriteFile(path, []byte(content), 0644)
	log.Log_err(err, "")
}
func mdIt(comment *domModel.Comment) string {
	md := ""
	md += Quote + BoldIt(comment.Name) + " said:" + EndLine
	md += Quote + comment.Content + EndSequence
	return md
}
func GetIt(url string) *domModel.CommentSection {
	path := file_pkg.Make_path([]string{DataDownPath}, url, ".md")
	if file_pkg.FileExists(path) {
		comments := domModel.CommentSection{
			Html: template.HTML(string(markml.MdToHtml(path))),
		}
		return &comments
	}
	return nil
}

func BoldIt(content string) string {
	return "**" + content + "**"
}
