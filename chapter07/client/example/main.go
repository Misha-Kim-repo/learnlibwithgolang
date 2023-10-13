package main

import "github.com/PacktPublishing/Go-Programming-Cookbook-Second-Edition/chapter7/client"

func main() {
	//Setup 함수의 첫 번쨰 인자는 true로 설정하고 두 번째 인자는 false로 설정한다.
	cli := client.Setup(true, false)
	if err := client.DefaultGetGolang(); err != nil {
		panic(err)
	}

	if err := client.DoOps(cli); err != nil {
		panic(err)
	}

	c := client.Controller{Client: cli}
	if err := c.DoOps(); err != nil {
		panic(err)
	}

	//Setup 함수의 첫 번째 인자를 true로 설정하고 두 번쨰 인자도 true로 설정한다.
	client.Setup(true, true)
	if err := client.DefaultGetGolang(); err != nil {
		panic(err)
	}
}
