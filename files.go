package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

const (
	ArticlesPath           = "Articles"
	GlobalAticlesDirectory = ArticlesPath + "/Glob"
	AttractionsDirectory   = ArticlesPath + "/Attractions"
	IntroDirectory         = ArticlesPath + "/Intro"
)

func MakePath(dir []string, file string, format string) string {
	Dir := ""
	for _, v := range dir {
		Dir += v + "/"
	}
	return Dir + file + "." + format
}
func ReadJson(path string, v interface{}) error {
	jsonFile, err := os.Open(path)
	if err != nil {
		return err
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, v)
	return nil
}

// func ReadArticle(name string) (*Article, error) {
// 	jsonFile, err := os.Open(AticlesDirectory + name + ".json")
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer jsonFile.Close()
// 	byteValue, _ := ioutil.ReadAll(jsonFile)
// 	var article Article
// 	json.Unmarshal(byteValue, &article)
// 	return &article, nil
// }

func ListDirectory(path string) ([]os.FileInfo, error) {
	return ioutil.ReadDir(path)

}

func LoadAttarctionsListFiles() *Attractions {
	var attractions Attractions
	files, err := ListDirectory(AttractionsDirectory + "/Docs")
	if nil != err {
		//HandleError
		return nil
	} else {
		for _, v := range files {
			var newAttarction Attraction
			path := MakePath([]string{AttractionsDirectory, "Docs"}, v.Name(), "")
			err := ReadJson(path, &newAttarction)
			if nil == err {
				attractions.Attractions = append(attractions.Attractions, newAttarction)
			}
		}
		attractions.Count = len(files)
		return &attractions
	}

}
func LoadIntroListFiles() *Articles {
	var doc Doc
	path := MakePath([]string{DocList}, "intro", "json")
	err := ReadJson(path, &doc)
	if nil != err {
		fmt.Println(err)
	}
	var intros []Article
	for _, v := range doc.Files {
		var file Article
		path := MakePath([]string{ArticlesPath, "Intro", "Docs"}, v, "json")
		err := ReadJson(path, &file)
		if nil != err {
			fmt.Println(err)
		}
		intros = append(intros, file)
	}
	articles := Articles{
		Title:    doc.Title,
		Articles: intros,
	}
	return &articles

}
