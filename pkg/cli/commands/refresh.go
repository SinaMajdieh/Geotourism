package commands

import (
	"github.com/SinaMajdieh/Geotourism/internal"
)

func InitRefreshCommand() Command {
	return Command{
		Flags: []Flag{},
		Run:   refresh,
	}
}

func refresh(args []string) {
	internal.Load_Data()
}
