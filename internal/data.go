package internal

import (
	"strings"

	"Geotourism/pkg/domModel"
	"Geotourism/pkg/markml"
	"Geotourism/pkg/tourson"
)

func Load_attractions() {
	attractionsList, attractions = tourson.Load_attractions_list_files(phenomena)
}
func Update_attractions_list(att *domModel.Attraction) {
	/*att.Value = strings.Replace(att.Phenomena, " ", "_", -1)
	attractionsList.Map[att.Value].Attractions = append(attractionsList.Map[att.Value].Attractions, *att)
	attractionsList.Map[att.Value].Count++

	attractions.Attractions = append(attractions.Attractions, *att)
	attractions.Count++*/
}
func Load_intros() {
	intros = tourson.LoadIntroListFiles(DocList)
}
func Load_blog() {
	blogPosts = markml.Read_posts()
}
func Load_Data() {
	Load_attractions()
	Load_intros()
	Load_blog()
}
func AllAttractions() *domModel.Attractions {
	return attractions
}
func contains(container, sub string) bool {
	if strings.Contains(strings.ToLower(container), strings.ToLower(sub)) {
		return true
	}
	return false
}
func Sattaction(index string) [](*domModel.Attraction) {
	result := make([](*domModel.Attraction), 0)
	for i, v := range attractions.Attractions {
		if contains(v.Title, index) {
			result = append(result, &attractions.Attractions[i])
			//fmt.Println(log.Highlight(v.Title, index))
			continue

		}
		f := false
		for _, l := range v.Content {
			if contains(l, index) {
				result = append(result, &attractions.Attractions[i])
				//fmt.Println(log.Highlight(l, index))
				f = true
				break
			}
		}
		if f {
			continue
		}
		for _, g := range v.Gallery {
			if contains(g.Caption, index) {
				result = append(result, &attractions.Attractions[i])
				//fmt.Println(log.Highlight(g.Caption, index))
				break
			}
		}
	}
	return result
}
