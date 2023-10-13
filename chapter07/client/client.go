package client

import (
	"crypto/tls"
	"net/http"
)

// Setup 함수는 클라이언트를 구성하고 전역 DefaultClient를 재정의한다.
func Setup(isSecure, nop bool) *http.Client {
	c := http.DefaultClient
	//때로는 테스트를 위해 SSL 인증 기능을 OFF 시켜야 할 때도 있다.
	if !isSecure {
		c.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: false,
			},
		}
	}
	if nop {
		c.Transport = &NopTransport{}
	}
	http.DefaultClient = c

	return c
}

// NopTransport는 아무 일도 하지 않는(No-Operation) 함수이다.
type NopTransport struct {
}

func (n *NopTransport) RoundTrip(*http.Request) (*http.Response, error) {
	//초기화되지 않은 응답이라는 것에 주의한다.
	//헤더 등을 살펴보면 알 수 있다.
	return &http.Response{StatusCode: http.StatusTeapot}, nil
}
