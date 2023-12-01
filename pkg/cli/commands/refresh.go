package commands

import (
	"fmt"

	"github.com/SinaMajdieh/Geotourism/internal"
)

func InitRefreshCommand() Command {
	return Command{
		Flags: []Flag{},
		Run:   refresh,
		Help:  HelpRefresh,
	}
}

func refresh(args []string) {
	internal.Load_Data()
}
func HelpRefresh() string {
	return fmt.Sprintf("%-40s%s", "refresh", "Reload the data\n")
}
