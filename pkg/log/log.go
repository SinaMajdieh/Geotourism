package log

import (
	"fmt"
	"github.com/TwiN/go-color"
)

func Log_err(err error, msg string) {
	if nil != err {
		fmt.Println(color.Ize(color.Red, err.Error()))
	} else {
		if 0 != len(msg) {
			fmt.Println(color.Ize(color.Green, msg))
		}
	}
}
