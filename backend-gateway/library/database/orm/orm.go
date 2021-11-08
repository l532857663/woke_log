package orm

import (
	logger "github.com/cihub/seelog"
	"github.com/jinzhu/gorm"
	// mysql 数据库驱动
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func NewDB(driver string, dsn string) *gorm.DB {
	db, err := gorm.Open(driver, dsn)
	if err != nil {
		logger.Error("db dsn(%s) error: %+v", dsn, err)
		panic(err)
	}

	// 不使用复数表名
	db.SingularTable(true)

	return db
}
