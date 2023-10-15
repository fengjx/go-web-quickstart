package db

import (
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/reflectx"

	"github.com/fengjx/go-web-quickstart/internal/app/appconfig"
)

var dbMap = make(map[string]*sqlx.DB)
var defaultDB *sqlx.DB

var toLowerMapper = reflectx.NewMapperFunc("json", strings.ToLower)

func Init() {
	for k, c := range appconfig.Conf.DB {
		db := sqlx.MustOpen(c.Type, c.Dsn)
		err := db.Ping()
		if err != nil {
			log.Panicf("db ping err - %s, %s, %s", c.Type, c.Dsn, err.Error())
		}
		if c.MaxIdle != 0 {
			db.SetMaxIdleConns(c.MaxIdle)
		}
		if c.MaxConn != 0 {
			db.SetMaxOpenConns(c.MaxConn)
		}
		db.Mapper = toLowerMapper
		dbMap[k] = db
	}
	defaultDB = dbMap["default"]
}
