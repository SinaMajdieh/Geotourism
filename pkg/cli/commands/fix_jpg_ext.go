package commands

import (
	"fmt"

	"Geotourism/internal"
	"Geotourism/pkg/tourson"
)

func InitFixJpgExtCommand() Command {
	return Command{
		Flags: []Flag{},
		Run:   fixJpgExt,
		Help:  HelpFixJpgExt,
	}
}
func fixJpgExt(args []string) {
	internal.Load_attractions()
	for _, att := range internal.AllAttractions().Attractions {
		for i, _ := range att.Gallery {
			if att.Gallery[i].Src[len(att.Gallery[i].Src)-4:] != ".jpg" {
				att.Gallery[i].Src = att.Gallery[i].Src + ".jpg"
			}
		}
		tourson.Save_attraction_file(&att)
	}
	internal.Load_attractions()
}
func HelpFixJpgExt() string {
	return fmt.Sprintf("%-40s%s", "fixJpgExt", "Fixing the jpg extension problem\n")
}
