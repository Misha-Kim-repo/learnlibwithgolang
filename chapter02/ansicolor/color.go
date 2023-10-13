package ansicolor

import "fmt"

//텍스트 색상
type Color int

const (
	ColorNone = iota

	//빨간색 테스트
	Red

	//초록색 테스트
	Green

	//노란색 테스트
	Yellow

	//파란색 테스트
	Blue

	//마젠타 테스트
	Magenta

	//청록색 테스트
	Cyan

	//흰색 테스트
	White

	//검정색 테스트
	Black Color = -1
)

//ColorText 구조체는 색상의 문자열 및 색상 값을 저장한다.
type ColorText struct {
	TextColor Color
	Text      string
}

func (r *ColorText) String() string {
	if r.TextColor == ColorNone {
		return r.Text
	}

	value := 30
	if r.TextColor != Black {
		value += int(r.TextColor)
	}
	return fmt.Sprintf("\033[0;%dm%s\033[0m", value, r.Text)
}
