package app

import (
	"database/sql"
	"fachrurozirizky/belajar-golang-restful-api/helper"
	"time"
)

func NewDB() *sql.DB {
	dataSource := "root:@tcp(localhost:3306)/belajar-golang-restful-api?parseTime=true"
	db, err :=sql.Open("mysql", dataSource)
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}