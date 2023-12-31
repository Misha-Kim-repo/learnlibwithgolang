package csvformat

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
)

// Movie 구조체는 파싱된(parsed) CSV를 저장하는 데 사용된다.
type Movie struct {
	Title    string
	Director string
	Year     int
}

// ReadCSV 함수는 io.Reader에 전달된 CSV를 처리하는 예제를 보여준다.
func ReadCSV(b io.Reader) ([]Movie, error) {
	r := csv.NewReader(b)
	//이 작업은 설정 옵션을 설정하는 선택하는 선택적 작업이다.
	r.Comma = ';'
	r.Comment = '-'
	var movies []Movie
	//지금은 header를 가져와 무시한다. header나 dictionary 키나 다른 형태의 룩업(lookup)으로 사용할 수도 있다.
	_, err := r.Read()
	if err != nil && err != io.EOF {
		return nil, err
	}

	//CSV를 모두 처리할 때까지 루프를 수행한다.
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		year, err := strconv.ParseInt(record[2], 10, 64)
		if err != nil {
			return nil, err
		}
		m := Movie{record[0], record[1], int(year)}
		movies = append(movies, m)
	}

	return movies, nil
}

// AddMoviesFromText 함수는 CSV 파서를 사용한다.
func AddMoviesFromText() error {
	in := `
- first our headers
movie title;director;year released

-then some data
Guadians of Galaxy vol. 2;James Gunn;2017
Star Wars: Episode VIII;Rian Johnson;2017
`
	b := bytes.NewBufferString(in)
	m, err := ReadCSV(b)
	if err != nil {
		return err
	}
	fmt.Printf("%#v\n", m)
	return nil
}
