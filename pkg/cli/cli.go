package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"Geotourism/pkg/cli/commands"
	"github.com/TwiN/go-color"
)

var cmds map[string]commands.Command

func readLine() []string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	return strings.Split(line, " ")
}
func initCommands() {
	cmds = make(map[string]commands.Command)
	cmds["refresh"] = commands.InitRefreshCommand()
	cmds["fixJpgExt"] = commands.InitFixJpgExtCommand()
	cmds["list_untracked_imgs"] = commands.InitListUntrackedImgs()
	cmds["lui"] = commands.InitListUntrackedImgs()
	cmds["simage"] = commands.InitSimage()
	cmds["sattaction"] = commands.InitSattaction()
	cmds["sat"] = commands.InitSattaction()
	cmds["help"] = InitHelpCommand()
	//cmds["shutdown"] = commands.InitShutdownCommand()
}
func Cmd() {
	initCommands()
	for {
		fmt.Print(color.Ize(color.Blue, "> "))
		args := readLine()
		if _, ok := cmds[args[0]]; !ok {
			fmt.Println(color.Ize(color.Red, "command not found!"))
		} else {
			cmds[args[0]].Run(args[1:])
		}
	}
}
