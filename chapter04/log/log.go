package log

import (
	"bytes"
	"fmt"
	"log"
)

// Log 함수는 로거(logger) 설정을 사용한다.
func Log() {
	//bytes.Buffer에 기록하기 위해 로거를 설정한다.
	buf := bytes.Buffer{}

	//두 번쨰 인자는 접두사이며, 마지막 인자는 논리연산자 or를 결합하여 설정하는 옵션이다.
	logger := log.New(&buf, "logger: ", log.Lshortfile|log.Ldate)
	logger.Println("test")
	logger.SetPrefix("New Logger: ")
	logger.Printf("You can also add args(%v) and use Fatalln to log and crash", true)
	fmt.Println(buf.String())
}
