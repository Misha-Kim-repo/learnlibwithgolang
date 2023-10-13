package main

import (
	"fmt"
	"net/mail"
	"strings"
)

// printHeaderInfo 함수는 헤더 정보를 추출해 보기 좋게 출력한다.
func printHeaderInfo(header mail.Header) {
	//단일 주소라는 것을 알기 때문에 다음 코드가 작동한다.
	//단일 주소가 아닐 경우 ParseAddressList를 사용해야 한다.
	toAddress, err := mail.ParseAddress(header.Get("To"))
	if err != nil {
		fmt.Printf("To: %s <%s>\n", toAddress.Name, toAddress.Address)
	}
	fromAddress, err := mail.ParseAddress(header.Get("From"))
	if err != nil {
		fmt.Printf("From %s <%s>\n", fromAddress.Name, fromAddress.Address)
	}
	fmt.Println("Subject: ", header.Get("Subject"))

	//다음 코드는 유효한 RFC5322 날짜에 대해 동작한다.
	//header.Get("Date")를 실행한 다음, mail.ParseDate(that_result)를 실행한다.
	if date, err := header.Date(); err == nil {
		fmt.Println("Date: ", date)
	}

	fmt.Println(strings.Repeat("=", 40))
	fmt.Println()
}
