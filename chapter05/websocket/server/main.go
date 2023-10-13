package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Listening on port: 8000")
	//모든 요청을 처리하기 위해 localhost:8000번 포트에 하나의 핸들러를 마운트한다.
	log.Panic(http.ListenAndServe("localhost:8000", http.HandlerFunc(wsHandler)))
}
