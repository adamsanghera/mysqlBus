package mysqlBus

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	address  = os.Getenv("MYSQL_ADDRESS")  // localhost default
	port     = os.Getenv("MYSQL_PORT")     // 3306 default
	username = os.Getenv("MYSQL_USERNAME") // root default
	password = os.Getenv("MYSQL_PASSWORD") // no pass default
	protocol = "tcp"
)

func init() {
	var err error
	DB, err = sql.Open("mysql", username+":"+password+"@"+protocol+"("+address+":"+port+")/app")
	if err != nil {
		panic(err)
	}
}
