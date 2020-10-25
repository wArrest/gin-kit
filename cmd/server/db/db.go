package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/wArrest/gin-kit/config"
	"os"
)

var db *sql.DB

func Init() {
	var err error
	db, err = sql.Open("mysql", config.GetConfig().Dsn)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func GetDB()*sql.DB  {
	return db
}
