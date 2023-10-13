package interfaces

import (
	"fmt"
	"io"
	"os"
)

// Copy 함수는 버퍼를 사용하여 in과 out의 첫번째 디렉터리로 데이터를 복사한다.
// 표준 출력(Stdout)에도 데이터를 쓴다.(write)
func Copy(in io.ReadSeeker, out io.Writer) error {
	//매개 변수 out과 표준 출력(Stdout)에도 데이터를 쓴다.
	w := io.MultiWriter(out, os.Stdout)
	//표준 복사, 매개 변수 in에 대량의 데이터가 있는 경우, 이 방법은 위험할 수 있다.
	if _, err := io.Copy(w, in); err != nil {
		return err
	}

	//64바이트의 청크(Chunk)를 사용하여 버퍼에 데이터를 쓴다.
	in.Seek(0, 0)
	buf := make([]byte, 64)

	if _, err := io.CopyBuffer(w, in, buf); err != nil {
		return err
	}

	//새 명령줄에 출력한다.
	fmt.Println()
	return nil
}
