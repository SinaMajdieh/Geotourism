package tourson

import (
	"fmt"
	"github.com/SinaMajdieh/Geotourism/pkg/domModel"
	"github.com/SinaMajdieh/Geotourism/pkg/file_pkg"
	"github.com/TwiN/go-color"
	"strconv"
	"strings"
)

// To log the LoadAttractionsListFiles function results
func log_attractions_loading_result(errors []string) {
	if 0 != len(errors) {
		fmt.Println(color.Ize(color.Yellow, strconv.Itoa(len(errors))+" error(s) found:"))
		for _, v := range errors {
			fmt.Println(color.Ize(color.Red, "Could not load "+v))
		}
	} else {
		fmt.Println(color.Ize(color.Green, "No errors detected. All is good!"))
	}
}

func LoadAttractionsListFiles(phenomena []string) (*domModel.AttractionsList, *domModel.Attractions) {
	var totattractions domModel.Attractions
	errored_files := make([]string, 0)
	attractions := make(map[string]*domModel.Attractions)
	for _, v := range phenomena {
		attractions[v] = &domModel.Attractions{}
		attractions[v].Phenomena = v
	}

	files, err := file_pkg.ListDirectory(file_pkg.AttractionsDirectory + "/docs")
	if nil != err {
		//HandleError
		return nil, nil
	} else {
		for _, v := range files {
			var newAttraction domModel.Attraction

			path := file_pkg.MakePath([]string{file_pkg.AttractionsDirectory, "docs"}, v.Name(), "")
			err := file_pkg.ReadJson(path, &newAttraction)
			if nil == err {
				newAttraction.Value = strings.Replace(newAttraction.Phenomena, " ", "_", -1)
				fmt.Println(color.Ize(color.Blue, "Successfully loaded "+newAttraction.Title))
				if _, ok := attractions[newAttraction.Value]; ok {
					attractions[newAttraction.Value].Attractions = append(attractions[newAttraction.Value].Attractions, newAttraction)
					attractions[newAttraction.Value].Count++
					totattractions.Attractions = append(totattractions.Attractions, newAttraction)
				} else {
					errored_files = append(errored_files, v.Name())
				}

			} else {
				errored_files = append(errored_files, v.Name())
			}
		}
		log_attractions_loading_result(errored_files)
		return &domModel.AttractionsList{
			Map: attractions,
		}, &totattractions
	}

}
func LoadIntroListFiles(DocList string) *domModel.Articles {
	var doc domModel.Doc
	path := file_pkg.MakePath([]string{DocList}, "intro", ".json")
	err := file_pkg.ReadJson(path, &doc)
	if nil != err {
		fmt.Println(err)
	}
	var intros []domModel.Article
	for _, v := range doc.Files {
		var file domModel.Article
		path := file_pkg.MakePath([]string{file_pkg.ArticlesPath, "intro", "docs"}, v, ".json")
		err := file_pkg.ReadJson(path, &file)
		if nil != err {
			fmt.Println(err)
		}
		intros = append(intros, file)
	}
	articles := domModel.Articles{
		Title:    doc.Title,
		Articles: intros,
	}
	return &articles

}
