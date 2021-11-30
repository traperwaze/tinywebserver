package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/traperwaze/tinywebserver/config"
)

var DB *sql.DB

func Init() {
	db, err := sql.Open("mysql", config.GetMysqlConnectionString())

	if err != nil{
		panic(err)
	}

	DB = db
}