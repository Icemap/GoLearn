package main

import (
	"log"
	"net/http"
	_"github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var DBEngine *xorm.Engine

func main() {
	DBEngine, _ = xorm.NewEngine("mysql", "root:@tcp(127.0.0.1:3306)/vistadb?charset=utf8")
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
