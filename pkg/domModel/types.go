package domModel

import (
	"encoding/json"
	"fmt"
	"html/template"
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
	//images     []string `json:"images"`
}
type Picture struct {
	Src     string `json:"src"`
	Caption string `json:"caption"`
}
type Articles struct {
	Title    string
	Articles []Article
}
type Attraction struct {
	Title          string    `json:"title" schema:"title"`
	Content        []string  `json:"content" schema:"content"`
	Image          string    `json:"image" schema:"image"`
	Link           string    `json:"link" schema:"link"`
	Accessibility  string    `json:"accessibility" schema:"accessibility"`
	SeasonsToVisit string    `json:"seasons" schema:"seasons"`
	Seasons        []string  `json:"seasons_arr" schema:"_"`
	Code           string    `json:"code" schema:"code"`
	Phenomena      string    `json:"phenomena" schema:"phenomena"`
	Value          string    `json:"_" schema:"_"`
	Province       string    `json:"province" schema:"province"`
	Significance   string    `json:"significance" schema:"significance"`
	MapImage       string    `json:"map-image" schema:"map-image"`
	Gallery        []Picture `json:"gallery" schema:"_"`
}

func (a Attraction) String() string {
	content := ""
	for _, l := range a.Content {
		content += l + "\n"
	}
	gallery := ""
	for _, p := range a.Gallery {
		gallery += "src : " + p.Src + "\n"
		gallery += "caption : " + p.Caption + "\n"
	}
	return fmt.Sprintf("Title:%v\nContent:\n%vImage:%v\nLink:%v\nAccessibility:%v\nSeasonsToVisit:%v\nCode:%v\nPhenomena:%v\nProvince:%v\nSignificance:%v\nGallery:\n%v",
		a.Title, content, a.Image, a.Link, a.Accessibility, a.SeasonsToVisit, a.Code, a.Phenomena, a.Province, a.Significance, gallery)
}

type Attractions struct {
	Attractions []Attraction
	Count       int
	Phenomena   string
}
type AttractionsList struct {
	Map map[string]*Attractions
}
type BlogPost struct {
	Title string
	Link  string
	Html  template.HTML
}
type Comment struct {
	Name    string `schema:"name"`
	Link    string
	Content string `schema:"comment"`
}
type CommentSection struct {
	Html template.HTML
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
