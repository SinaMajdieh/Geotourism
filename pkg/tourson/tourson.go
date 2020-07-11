package tourson

import (
	"fmt"
	"github.com/SinaMajdieh/Geotourism/pkg/domModel"
	"github.com/SinaMajdieh/Geotourism/pkg/file_pkg"
	"strings"
)
func LoadAttractionsListFiles(phenomena []string) (*domModel.AttractionsList , *domModel.Attractions) {
	var totattractions domModel.Attractions
	attractions := make(map[string]*domModel.Attractions)
	for _ , v := range phenomena{
		attractions[v] = &domModel.Attractions{}
		attractions[v].Phenomena = v
	}

	files, err := file_pkg.ListDirectory(file_pkg.AttractionsDirectory + "/docs")
	if nil != err {
		//HandleError
		return nil,nil
	} else {
		for _, v := range files {
			var newAttraction domModel.Attraction

			path := file_pkg.MakePath([]string{file_pkg.AttractionsDirectory, "docs"}, v.Name(), "")
			err := file_pkg.ReadJson(path, &newAttraction)
			if nil == err {
				newAttraction.Value = strings.Replace(newAttraction.Phenomena , " " , "_" , -1)
				fmt.Println(newAttraction.Title)
				attractions[newAttraction.Value].Attractions = append(attractions[newAttraction.Value].Attractions , newAttraction)
				attractions[newAttraction.Value].Count++
				totattractions.Attractions = append(totattractions.Attractions, newAttraction)
			}
		}
		return &domModel.AttractionsList{
			Map : attractions,
		} , &totattractions
	}

}
func LoadIntroListFiles(DocList string) *domModel.Articles {
	var doc domModel.Doc
	path := file_pkg.MakePath([]string{DocList}, "intro", "json")
	err := file_pkg.ReadJson(path, &doc)
	if nil != err {
		fmt.Println(err)
	}
	var intros []domModel.Article
	for _, v := range doc.Files {
		var file domModel.Article
		path := file_pkg.MakePath([]string{file_pkg.ArticlesPath, "intro", "docs"}, v, "json")
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