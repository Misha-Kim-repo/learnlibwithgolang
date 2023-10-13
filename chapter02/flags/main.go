package main

import (
	"flag"
	"fmt"
)

func main() {
	c := Config{}
	c.Setup()

	flag.Parse()

	fmt.Print(c.GetMessage())
}
