package dao

import (
	"github.com/go-xorm/xorm"
	_"github.com/go-sql-driver/mysql"
)

var DBEngine *xorm.Engine

func init()  {
	DBEngine, _ = xorm.NewEngine("mysql", "root:@tcp(127.0.0.1:3306)/vistadb?charset=utf8")
}
