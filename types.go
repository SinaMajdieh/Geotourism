package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Doc struct {
	Title  string   `json:"title"`
	Link   string   `json:"link"`
	Docs   string   `json:"docs"`
	Images string   `json:"images"`
	Files  []string `json:"files"`
	Format string   `json:"format"`
}
type Article struct {
	Title   string   `json:"title"`
	Image   string   `json:"image"`
	Content []string `json:"content"`
	//Images     []string `json:"Images"`
}
type Articles struct {
	Title    string
	Articles []Article
}
type Attraction struct {
	Title          string   `json:"title"`
	Content        []string `json:"content"`
	Image          string   `json:"image"`
	Link           string   `json:"link"`
	Accessiblity   string   `json:"accessiblity"`
	SeasonsToVisit string   `json:"seasons"`
	Code           string   `json:"code"`
}
type Attractions struct {
	Attractions []Attraction
	Count       int
}

func (attraction *Attraction) ReadFile(path string) error {
	jsonFile, err := os.Open(path)
	if err != nil {
		return err
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, attraction)
	return nil
}
