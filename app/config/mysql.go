package config

import (
	"os"
)

type mysql struct {
	User      string
	Password  string
	Host      string
	Port      string
	DB        string
	Collation string
	Charset   string
}

var insMysql *mysql

func InsMysql() *mysql {
	if insMysql == nil {
		insMysql = &mysql{}
		insMysql.User = os.Getenv("MYSQL_USER")
		insMysql.Password = os.Getenv("MYSQL_PASSWORD")
		insMysql.Host = os.Getenv("MYSQL_HOST")
		insMysql.Port = os.Getenv("MYSQL_PORT")
		insMysql.DB = os.Getenv("MYSQL_DB")
		insMysql.Collation = os.Getenv("MYSQL_COLLATION")
		insMysql.Charset = os.Getenv("MYSQL_CHARSET")
	}
	return insMysql
}
