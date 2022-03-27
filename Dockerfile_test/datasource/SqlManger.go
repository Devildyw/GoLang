package datasource

import (

	"database/sql"
	"log"
	_"github.com/go-sql-driver/mysql"
)


//case-insensitive file name collision: "Sqlconn.go" and "sqlconn.go"go


func DataSourceInit() (*sql.DB) {
	db, err := sql.Open("mysql", "ding:dyw20020304@tcp(81.68.186.20:3306)/Ding?charset=utf8&parseTime=True")
	if err != nil {
		log.Println("sql.Open err = ", err)
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)
	db.SetConnMaxLifetime(20)
	//defer db.Close()
	err2 := db.Ping()
	if err2!=nil{
		log.Println(err2)
	}
	return db
}
