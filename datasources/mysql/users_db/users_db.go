package users_db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	username = "root"
	password = "6lNxw2NvHA"
	host     = "127.0.0.1:3306"
	schema   = "users_db"
)

var (
	Client *sql.DB

	/*
		# alternative get from environment variables
		username = os.Getenv(mysql_username)
		password = os.Getenv(mysql_password)
		host     = os.Getenv(mysql_host)
		schema   = os.Getenv(mysql_schema)
	*/
)

func init() {
	datasourceName := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8",
		username,
		password,
		host,
		schema)

	var err error

	Client, err = sql.Open("mysql", datasourceName)

	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("DB Connection success!")

}
