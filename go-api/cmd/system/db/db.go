package db

import (
	// from xorm.io
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

// Connect ORM Engine
func Connect() (db *xorm.Engine, err error) {
	/* db, err = xorm.NewEngine("mysql", "test:test_golang@/test?charset=utf8")
	if err != nil {
		return
	}
	db.DBMetas()
	return
	*/ // testing connection
	return xorm.NewEngine("mysql", "test:test_golang@/test?charset=utf8")
}
