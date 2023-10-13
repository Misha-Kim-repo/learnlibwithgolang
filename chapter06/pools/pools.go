package pools

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// Setup 함수는 pools 패키지를 사용하고 연결 수 제한 등을 적용해 db를 구성한다.
func Setup() (*sql.DB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/gocookbook?parseTime=true", os.Getenv("MYSQLUSERNAME"), os.Getenv("MYSQLPASSWORD")))
	if err != nil {
		return nil, err
	}

	//최대 24개의 연결만 허용한다.
	db.SetMaxOpenConns(24)

	//MaxIdleConns는 열려 있는 최대 SetMaxOpenConns 보다 작을 수 없다. 그렇지 않으면 해당 값으로 기본 설정된다.
	db.SetConnMaxIdleTime(24)

	return db, nil
}
