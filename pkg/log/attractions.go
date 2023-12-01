package log

import (
	"fmt"

	"Geotourism/pkg/domModel"
)

func HighlightSattaction(a *domModel.Attraction, index string) string {
	result := ""
	c := *a
	c.Title = Highlight(a.Title, index)
	content := make([]string, 0)
	for _, l := range a.Content {
		content = append(content, Highlight(l, index))
	}
	c.Content = content
	gallery := make([]domModel.Picture, 0)
	for _, g := range a.Gallery {
		p := domModel.Picture{}
		p.Src = g.Src
		p.Caption = Highlight(g.Caption, index)
		gallery = append(gallery, p)
	}
	c.Gallery = gallery
	result = fmt.Sprintln(c)
	return result
}
