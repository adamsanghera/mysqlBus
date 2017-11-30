package mysqlBus

import (
	"database/sql"
	// hehe
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	var err error
	DB, err = sql.Open("mysql", "root@tcp(database:3306)/app")
	if err != nil {
		panic(err)
	}
}
