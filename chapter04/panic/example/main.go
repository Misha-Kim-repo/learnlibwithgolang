package main

import (
	"fmt"

	"github.com/PacktPublishing/Go-Cookbook-Second-Edition/chapter4/panic"
)

func main() {
	fmt.Println("before panic")
	panic.Catcher()
	fmt.Println("after panic")
}
