package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" //we import supported libraries for database/sql
)

// Query 함수는 새로운 연결을 가져와 테이블을 생성하고 일부 쿼리 작업을 한 다음, 나중에 테이블을 삭제한다.
func Query(db *sql.DB, name string) error {
	rows, err := db.Query("SELECT name, created FROM example where name=?", name)
	if err != nil {
		return err
	}
	defer db.Close()
	for rows.Next() {
		var e Example
		if err := rows.Scan(&e.Name, &e.Created); err != nil {
			return err
		}
		fmt.Printf("Result: \n\tName: %s\n\tCreated: %v\n", e.Name, e.Created)
	}
	return rows.Err()
}
