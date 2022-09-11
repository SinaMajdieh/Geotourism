package commands

import (
	"fmt"
	"github.com/SinaMajdieh/Geotourism/internal"
	"github.com/SinaMajdieh/Geotourism/pkg/file_pkg"
	"github.com/TwiN/go-color"
)

const (
	prefixLen = len("/docimg/Attractions/")
	imgsDir   = "./website/articles/img/Attractions/"
)

func InitListUntrackedImgs() Command {
	return Command{
		Flags: []Flag{},
		Run:   listUntrackedImgs,
	}
}
func listUntrackedImgs(args []string) {
	list := make([]string, 0)
	for _, att := range internal.AllAttractions().Attractions {
		for _, pic := range att.Gallery {
			name := pic.Src[prefixLen:]
			path := imgsDir + name
			if !file_pkg.FileExists(path) {
				list = append(list, att.Title+" : "+name)
			}
		}
	}
	for _, name := range list {
		fmt.Println(color.Ize(color.Yellow, name))
	}
	fmt.Println(color.Ize(color.Red, fmt.Sprintf("found %d untracked images.", len(list))))
}
