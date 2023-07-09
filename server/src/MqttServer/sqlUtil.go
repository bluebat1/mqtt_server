package MqttServer

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var (
	SqlUtil *sql.DB
)

func SqlUtilInit() error {
	var err error
	SqlUtil, err = sql.Open("mysql", ConfigGet("mysql.user")+":"+ConfigGet("mysql.passwd")+"@tcp("+ConfigGet("mysql.ip")+":"+ConfigGet("mysql.port")+")/"+ConfigGet("mysql.database"))
	if err != nil {
		return err
	}

	return nil
}
