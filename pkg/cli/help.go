package cli

import (
	"fmt"

	"github.com/SinaMajdieh/Geotourism/pkg/cli/commands"
	"github.com/TwiN/go-color"
)

func InitHelpCommand() commands.Command {
	return commands.Command{
		Flags: []commands.Flag{},
		Run:   Help,
		Help:  HelpHelp,
	}
}
func Help(args []string) {
	if len(args) == 0 {
		for _, c := range cmds {
			fmt.Print(color.Ize(color.Cyan, c.Help()))
		}
	} else {
		if c, ok := cmds[args[0]]; ok {
			fmt.Print(color.Ize(color.Cyan, c.Help()))
		} else {
			fmt.Println(color.Ize(color.Red, "command not found"))
		}
	}
}
func HelpHelp() string {
	return fmt.Sprintf("%-40s%s", "help (index)", "Help for other commands\n")
}
