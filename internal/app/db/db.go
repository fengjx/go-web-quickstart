package db

import "xorm.io/xorm"

func Default() *xorm.Engine {
	return defaultEngine
}

func GetConn(name string) *xorm.Engine {
	return engineMap[name]
}
