package commands

import (
	"fmt"
	"strings"

	"Geotourism/pkg/file_pkg"
	"Geotourism/pkg/log"
	"github.com/TwiN/go-color"
)

func InitSimage() Command {
	return Command{
		Flags: []Flag{},
		Run:   simage,
		Help:  HelpSimage,
	}
}

const imgPrefix = "website/articles/img/"

func simage(args []string) {
	if len(args) < 2 {
		fmt.Println(color.Ize(color.Red, "simage -flag(-e for equal, -c for contain) 'index'"))
	} else {
		index := args[1]
		dirs, err := file_pkg.ListDirectory(imgPrefix)
		log.Log_err(err, "read image directory successfully")
		for _, dir := range dirs {
			if dir.IsDir() {
				fmt.Println(dir.Name() + "->")
				files, err := file_pkg.ListDirectory(imgPrefix + dir.Name() + "/")
				log.Log_err(err, "")
				for _, file := range files {
					switch args[0] {
					case "-e":
						if file_pkg.RemoveFileExt(strings.ToLower(file.Name())) == strings.ToLower(index) {
							fmt.Println(color.Ize(color.Green, "\t"+file.Name()))
						} else {
							fmt.Println(color.Ize(color.White, "\t"+file.Name()))
						}
					case "-c":
						if strings.Contains(file_pkg.RemoveFileExt(strings.ToLower(file.Name())), strings.ToLower(index)) {
							fmt.Println("\t" + log.Highlight(file.Name(), index))
						} else {
							fmt.Println(color.Ize(color.White, "\t"+file.Name()))
						}
					default:
						break

					}

				}
			}
		}

	}
}
func HelpSimage() string {
	return fmt.Sprintf("%-40s%s", "simage (-c:contains -e:equals) (index)", "Search for images\n")
}
