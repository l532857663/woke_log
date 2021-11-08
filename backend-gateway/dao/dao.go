package dao

import (
	"backend-gateway/conf"
	"backend-gateway/dao/cache"
	"backend-gateway/dao/db"

	"github.com/jinzhu/gorm"
)

const (
	DBNAME_BACKEND = "backend"
)

var d *Dao // DAO 层全局变量

// @Description DAO管理资源结构
// @Author Wangch
// @Version 1.0
// @Update Wangch 2021-09-26 init
type Dao struct {
	Conf  *conf.Config // 配置对象
	db    *db.DB       // 数据库操作相关对象
	cache *cache.Cache // 缓存操作相关对象
}

// @Description 创建一个 DAO 并返回对象
// @Author Wangch
// @Version 1.0
// @Update Wangch 2021-09-26 init
func New(c *conf.Config, appletName string) Storage {
	if d == nil {
		d = &Dao{
			Conf: c,
		}

		// 创建数据库连接对象
		d.db = db.New(c, appletName)

		// 创建缓存连接对象
		if c.Cache.Enabled {
			d.cache = cache.New(c.Cache)
		}
	}
	return d
}

// @Description 关闭DAO层创建的资源
// @Author Wangch
// @Version 1.0
// @Update Wangch 2021-09-26 init
func (d *Dao) Close() {
	d.db.Close()

	if d.Conf.Cache.Enabled {
		d.cache.Close()
	}
}

// @Description 根据 model 同步指定数据库
// @Author Ryen
// @Version 1.0
// @Update Ryen 2020-04-27 init
func (d *Dao) SyncDB(dbName string) {
	d.db.SyncDB(dbName)

	return
}

// @Description 事务回滚控制（出错返回true）
// @Author Wangch
// @Version 1.0
// @Update Wangch 2021-09-26 init
func DbUpdateResultRollbackControl(db *gorm.DB) bool {
	if db.RowsAffected == 0 && db.Error != nil {
		return true
	}

	return false
}

// NOTE: 添加数据库，补充该方法
// @Description 获取数据库事务
// @Author Wangch
// @Version 1.0
// @Update Wangch 2021-09-26 init
func (d *Dao) GetDbTxn(dbName string) *gorm.DB {
	var txn *gorm.DB

	switch dbName {
	case DBNAME_BACKEND:
		txn = d.db.DBBackend.Begin()
	}

	return txn
}

// NOTE: 添加数据库，补充该方法
// @Description 获取数据库当前连接
// @Author Wangch
// @Version 1.0
// @Update Wangch 2021-09-26 init
func (d *Dao) GetDbConn(dbName string) *gorm.DB {
	switch dbName {
	case DBNAME_BACKEND:
		return d.db.DBBackend
	}

	return nil
}
