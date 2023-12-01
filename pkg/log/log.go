package log

import (
	"fmt"
	"strings"

	"github.com/TwiN/go-color"
)

const (
	multiplyMax = 100
)

func Error(msg string, from string) {
	fmt.Println(color.Ize(color.Red, msg), " from:", color.Ize(color.Yellow, from))
}
func Log_err(err error, msg string) {
	if nil != err {
		fmt.Println(color.Ize(color.Red, err.Error()))
	} else {
		if 0 != len(msg) {
			fmt.Println(color.Ize(color.Green, msg))
		}
	}
}
func findOccurrence(container, sub string) []int {
	occ := make([]int, 0)
	container = strings.ToLower(container)
	sub = strings.ToLower(sub)
	for i := range container {
		if container[i] == sub[0] {
			ok := 0
			for j := 0; j < len(sub) && i+j < len(container); j++ {
				if container[i+j] != sub[j] {
					break
				} else {
					ok++
				}
			}
			if ok == len(sub) {
				occ = append(occ, i)
				i += len(sub)
			}
		}
	}
	return occ
}
func Highlight(container, sub string) string {
	result := ""
	occ := findOccurrence(container, sub)
	prev := 0
	for _, x := range occ {
		result += fmt.Sprint(container[prev:x])
		result += fmt.Sprint(color.Ize(color.Green, container[x:x+len(sub)]))
		prev = x + len(sub)
	}
	if prev < len(container) {
		result += fmt.Sprint(container[prev:])
	}
	return result
}
func Multiply(s string, t int) string {
	r := ""
	if t > multiplyMax {
		Error("Too long!", "Log.Multiply()")
		return s
	} else {
		for ; t > 0; t-- {
			r += s
		}
	}
	return r
}
