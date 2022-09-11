package file_pkg

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	ArticlesPath            = "website/articles"
	GlobalArticlesDirectory = ArticlesPath + "/glob"
	AttractionsDirectory    = ArticlesPath + "/attractions"
	IntroDirectory          = ArticlesPath + "/intro"
)

func Make_path(dir []string, file string, format string) string {
	Dir := ""
	for _, v := range dir {
		Dir += v + "/"
	}
	return Dir + file + format
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
// 	jsonFile, err := os.Open(AticlesDirectory + name + "Chele' Gah.json")
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

// FileExists check if a file exists
func FileExists(path string) bool {
	if _, err := os.Stat(path); nil != err {
		return false
	}
	return true
}

// RemoveFileExt removes file name's extension
func RemoveFileExt(fileName string) string {
	return fileName[:len(fileName)-len(filepath.Ext(fileName))]
}

// AddFileExt adds an extension to a file's name
func AddFileExt(fileName, ext string) string {
	return fileName + "." + ext
}
