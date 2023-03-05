package db

import (
	"fengjx/go-web-quickstart/internal/app/appconfig"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

var engineMap = make(map[string]*xorm.Engine)
var defaultEngine *xorm.Engine

func Init() {
	for k, c := range appconfig.Conf.DB {
		e, err := xorm.NewEngine(c.Type, c.Dsn)
		if err != nil {
			log.Panicf("create db connection err - %s, %s, %s", c.Type, c.Dsn, err.Error())
		}
		e.Omit("ctime", "utime")
		e.ShowSQL(c.ShowSQL)
		if c.MaxIdle != 0 {
			e.SetMaxIdleConns(c.MaxIdle)
		}
		if c.MaxConn != 0 {
			e.SetMaxOpenConns(c.MaxConn)
		}
		e.SetMapper(names.GonicMapper{})
		engineMap[k] = e
	}
	defaultEngine = engineMap["default"]
}
