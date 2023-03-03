package db

import (
	"fengjx/go-web-quickstart/internal/app/appconfig"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"xorm.io/xorm"
)

var engineMap map[string]*xorm.Engine

var defaultEngine *xorm.Engine

func Init() {
	for k, c := range appconfig.Conf.DB {
		e, err := xorm.NewEngine(c.Type, c.Dsn)
		if err != nil {
			log.Panicf("create db connection err - %s, %s, %s", c.Type, c.Dsn, err.Error())
		}
		engineMap[k] = e
	}
	defaultEngine = engineMap["default"]
}
