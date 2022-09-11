package cli

import (
	"bufio"
	"fmt"
	"github.com/SinaMajdieh/Geotourism/pkg/cli/commands"
	"github.com/TwiN/go-color"
	"os"
	"strings"
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
	//cmds["shutdown"] = commands.InitShutdownCommand()
}
func Cmd() {
	initCommands()
	for {
		args := readLine()
		if _, ok := cmds[args[0]]; !ok {
			fmt.Println(color.Ize(color.Red, "command not found!"))
		} else {
			cmds[args[0]].Run(args[1:])
		}
	}
}
