package cmd

import (
	"fmt"
	"github.com/gookit/color"
)

var red = color.FgRed.Render
var white = color.FgWhite.Render
var green = color.FgGreen.Render
var yellow = color.FgYellow.Render

func Error(label string) string {
	bracket := fmt.Sprintf("%s%s%s", red("["), white("!"), red("]"))
	return fmt.Sprintf("%s %s", bracket, red(label))
}

func PrintInfo(label string, value string) {
	fmt.Printf("%s: %s\n", green(label), yellow(value))
}

func PrintNumber(label string, value int, suffix string) {
	fmt.Printf("%s: %s %s\n", green(label), yellow(value), green(suffix))
}
