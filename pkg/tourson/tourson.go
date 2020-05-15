package tourson

import (
	"fmt"
	"github.com/SinaMajdieh/Geotourism/pkg/domModel"
	"github.com/SinaMajdieh/Geotourism/pkg/filePkg"
)
func LoadAttractionsListFiles(phenomena []string) (*domModel.AttractionsList , *domModel.Attractions) {
	var totattractions domModel.Attractions

	var attractionsErosion domModel.Attractions
	var attractionsSedimentary domModel.Attractions
	files, err := filePkg.ListDirectory(filePkg.AttractionsDirectory + "/Docs")
	if nil != err {
		//HandleError
		return nil,nil
	} else {
		for _, v := range files {
			var newAttraction domModel.Attraction

			path := filePkg.MakePath([]string{filePkg.AttractionsDirectory, "Docs"}, v.Name(), "")
			err := filePkg.ReadJson(path, &newAttraction)
			if nil == err {
				if newAttraction.Phenomena == phenomena[1] {
					attractionsErosion.Attractions = append(attractionsErosion.Attractions, newAttraction)
					attractionsErosion.Count++
				}else if newAttraction.Phenomena == phenomena[2]{
					attractionsSedimentary.Attractions = append(attractionsSedimentary.Attractions , newAttraction)
					attractionsSedimentary.Count++
				}

				totattractions.Attractions = append(totattractions.Attractions, newAttraction)

			}
		}

		return &domModel.AttractionsList{
			Erosion: attractionsErosion,
			Sedimentary:attractionsSedimentary,
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
		path := filePkg.MakePath([]string{filePkg.ArticlesPath, "Intro", "Docs"}, v, "json")
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