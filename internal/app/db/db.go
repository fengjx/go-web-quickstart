package db

import (
	"github.com/jmoiron/sqlx"
)

func GetDefaultDB() *sqlx.DB {
	return defaultDB
}

func GetDB(name string) *sqlx.DB {
	return dbMap[name]
}
