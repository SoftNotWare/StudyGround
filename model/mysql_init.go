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
	db, err = sql.Open("mysql", "root:root@tcp(10.0.0.3)/study_ground?timeout=3s&writeTimeout=5s&readTimeout=2s&charset=utf8mb4&parseTime=true&loc=Local")
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Fatal("连接数据库失败")
	}

	err = db.Ping()
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Fatal("ping数据库失败")
	}
}
