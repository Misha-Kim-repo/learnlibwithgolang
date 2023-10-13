package panic

import (
	"fmt"
	"strconv"
)

// Paic 함수는 분모를 0으로 취하는 것으로 인해 패닉을 발생한다.
func Panic() {
	zero, err := strconv.ParseInt("0", 10, 64)
	if err != nil {
		panic(err)
	}
	a := 1 / zero
	fmt.Println("We'll never get here", a)
}

// Catcher 함수는 Painc 함수를 호출한다.
func Catcher() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic occurred: ", r)
		}
	}()
	Panic()
}
