package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// CatchSig 함수는 SIGINT 인터럽트에 대한 리스너를 설정한다.
func CatchSig(ch chan os.Signal, done chan bool) {
	//신호에 대기하기 위해 블록(block)한다.
	sig := <-ch
	//신호를 받으면 출력한다.
	fmt.Println("\nsig received: ", sig)

	//모든 신호에 대한 핸들러(handler)를 설정할 수 있다.
	switch sig {
	case syscall.SIGINT:
		fmt.Println("handling a SIGINT now!")
	case syscall.SIGTERM:
		fmt.Println("handling a SIGTERM in an entirely different way!")
	default:
		fmt.Println("unexpected signal received")
	}
	done <- true
}

func main() {
	//채널을 초기화한다.
	signals := make(chan os.Signal)
	done := make(chan bool)

	//signals 라이브러리에 연결한다.
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	//이 GO 루틴에 의해 신호가 잡힐 경우 done으로 쓴다.
	go CatchSig(signals, done)
	fmt.Println("Press ctrl-c to terminate...")

	//누군가 done으로 쓸 때까지 프로그램을 블로킹한다.
	<-done
	fmt.Println("Done !")
}
