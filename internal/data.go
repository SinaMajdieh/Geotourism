package internal

import (
	"github.com/SinaMajdieh/Geotourism/pkg/domModel"
	"github.com/SinaMajdieh/Geotourism/pkg/tourson"
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
func Load_Data() {
	Load_attractions()
	Load_intros()
}
