package commands

import (
	"fmt"
	"strconv"

	"Geotourism/internal"
	"Geotourism/pkg/log"
	"github.com/TwiN/go-color"
)

func InitSattaction() Command {
	return Command{
		Flags: []Flag{},
		Run:   Sattaction,
		Help:  HelpSattaction,
	}
}

func Sattaction(args []string) {
	searchResult := internal.Sattaction(args[0])
	fmt.Println(log.Multiply("=", 80))
	for _, v := range searchResult {
		fmt.Println(log.HighlightSattaction(v, args[0]))
		fmt.Println(log.Multiply("=", 80))
	}
	if x := len(searchResult); x > 0 {
		fmt.Println(color.Ize(color.Green, strconv.Itoa(x)+" attraction(s) found!"))
	} else {
		fmt.Println(color.Ize(color.Red, "No attractions found!"))
	}
}
func HelpSattaction() string {
	return fmt.Sprintf("%-40s%s", "sat (index)", "Search the attractions for index\n")
}
