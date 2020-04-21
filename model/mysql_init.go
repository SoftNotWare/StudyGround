package model

import (
	"database/sql"

	// _ 引入mysql
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

var (
	db *sql.DB
)

// InitDB 初始化数据库
func InitDB() {
	var err error
	db, err = sql.Open("mysql", Config.DSN)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Fatal("连接数据库失败")
	}

	err = db.Ping()
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Fatal("ping数据库失败")
	}
}
