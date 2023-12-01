package commands

import (
	"fmt"

	"Geotourism/internal"
	"Geotourism/pkg/file_pkg"
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
		Help:  HelplistUntrackedImgs,
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
func HelplistUntrackedImgs() string {
	return fmt.Sprintf("%-40s%s", "list_untracked_imgs", "List the missing images\n")
}
