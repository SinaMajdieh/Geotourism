package main

func FindAttraction(link string) *Attraction {
	for _, v := range attractions.Attractions {
		if v.Link == link {
			return &v
		}
	}
	return nil
}
