package main

import "guthub.com/PacktPublishing/Go-Programming-Cookbook-Second-Edition/global"

func main() {
	if err := global.UseLog(); err != nil {
		panic(err)
	}
}
