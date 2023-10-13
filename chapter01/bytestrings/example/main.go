package main

import "github.com/PacktPublishing/Go-Programming-Cookbook-Second-Edition/chapter1/bytestrings"

func main() {
	err := bytestrings.WorkWithBuffer()
	if err != nil {
		panic(err)
	}

	//아래 내용은 모두 표준 출력(Stdout)에 출력된다.
	bytestrings.SearchString()
	bytestrings.ModifyString()
	bytestrings.StringReader()
}
