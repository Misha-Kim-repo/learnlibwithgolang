package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/gorilla/websocket"
)

// catSig 함수는 ctrl-c로 프로그램을 종료하면 웹소켓 연결을 정리한다.
func catchSig(ch chan os.Signal, c *websocket.Conn) {
	//os.Signal을 대기하기 위해 블록한다.
	<-ch
	err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		log.Println("Write Close:", err)
	}
	return
}

func main() {
	//채널에 os.Signal 값을 연결한다.
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	//웹소켓에 연결하기 위해 ws:// 구조를 사용한다.
	u := "ws://localhost:8000/"
	log.Printf("connecting to %s", u)

	c, _, err := websocket.DefaultDialer.Dial(u, nil)
	if err != nil {
		log.Fatal("dial: ", err)
	}
	defer c.Close()

	//catchSig에 전달한다.
	go catchSig(interrupt, c)

	process(c)
}
