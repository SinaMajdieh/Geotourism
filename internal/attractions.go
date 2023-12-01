package internal

import (

	"Geotourism/pkg/domModel"
)

func FindAttraction(link string) *domModel.Attraction {
	for _, v := range attractions.Attractions {
		if v.Link == link {
			return &v
		}
	}
	return nil
}