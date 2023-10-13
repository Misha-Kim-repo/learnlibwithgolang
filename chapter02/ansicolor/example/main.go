package main

import (
	"fmt"

	"github.com/PacktPublishing/Go-Programming-Second-Edition/chapter2/ansicolor"
)

func main() {
	r := ansicolor.ColorText{
		TextColor: ansicolor.Red,
		Text:      "I'm Red!",
	}
	fmt.Println(r.String())

	r.TextColor = ansicolor.Green
	r.Text = "Now I'm Green!"
	fmt.Println(r.String())

	r.TextColor = ansicolor.ColorNone
	r.Text = "Back to normal..."
	fmt.Println(r.String())
}
