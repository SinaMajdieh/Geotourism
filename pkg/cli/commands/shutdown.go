package commands

import (
	"context"
	"fmt"
	"github.com/TwiN/go-color"
)

func InitShutdownCommand() Command {
	return Command{
		Flags: []Flag{},
		Run:   shutDown,
	}
}
func shutDown(args []string) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	<-ctx.Done()
	fmt.Println(color.Ize(color.Green, "trying to shutdown server..."))
}
