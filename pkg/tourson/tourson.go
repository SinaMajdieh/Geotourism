package tourson

import (
	"fmt"
	"github.com/SinaMajdieh/Geotourism/pkg/domModel"
	"github.com/SinaMajdieh/Geotourism/pkg/filePkg"
)
func LoadAttractionsListFiles(phenomena []string) (*domModel.AttractionsList , *domModel.Attractions) {
	var totattractions domModel.Attractions
	attractions := make(map[string]*domModel.Attractions)
	for _ , v := range phenomena{
		attractions[v] = &domModel.Attractions{}
		attractions[v].Phenomena = v
	}

	files, err := filePkg.ListDirectory(filePkg.AttractionsDirectory + "/docs")
	if nil != err {
		//HandleError
		return nil,nil
	} else {
		for _, v := range files {
			var newAttraction domModel.Attraction

			path := filePkg.MakePath([]string{filePkg.AttractionsDirectory, "docs"}, v.Name(), "")
			err := filePkg.ReadJson(path, &newAttraction)
			if nil == err {
				attractions[newAttraction.Phenomena].Attractions = append(attractions[newAttraction.Phenomena].Attractions , newAttraction)
				attractions[newAttraction.Phenomena].Count++
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
	path := filePkg.MakePath([]string{DocList}, "intro", "json")
	err := filePkg.ReadJson(path, &doc)
	if nil != err {
		fmt.Println(err)
	}
	var intros []domModel.Article
	for _, v := range doc.Files {
		var file domModel.Article
		path := filePkg.MakePath([]string{filePkg.ArticlesPath, "intro", "docs"}, v, "json")
		err := filePkg.ReadJson(path, &file)
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