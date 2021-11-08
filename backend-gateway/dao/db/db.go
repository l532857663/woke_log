package db

import (
	"backend-gateway/conf"
	"backend-gateway/library/config"
	"backend-gateway/library/database/orm"
	"backend-gateway/model"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

// @Description DAO DB管理资源结构
// @Author Wangch
// @Version 1.0
// @Update Wangch 2021-09-26 init
type DB struct {
	DBBackend *gorm.DB // Backend 信息库
}

// @Description 创建一个 DAO 并返回对象
// @Author Wangch
// @Version 1.0
// @Update Wangch 2021-09-26 init
func New(c *conf.Config, appletName string) *DB {
	d := &DB{}

	// 创建数据库连接对象
	d.DBBackend = orm.NewDB("mysql", c.DBBackend.Dsn)

	d.setDbLogMode(c, appletName)

	return d
}

// NOTE: 添加数据库，补充该方法
// @Description 设置DAO层数据源日志模式
// @Author Wangch
// @Version 1.0
// @Update Wangch 2021-09-26 init
func (d *DB) setDbLogMode(c *conf.Config, appletName string) {
	var flag bool = false
	switch c.ServerGeneral.RunEnv {
	case config.RUN_ENV_PROD, config.RUN_ENV_BETA:
		flag = false
	default:
		flag = true
	}

	d.DBBackend.LogMode(flag)
	d.DBBackend.Set("gorm:table_options", "charset=utf8mb4")
	// 设置最大连接数、空闲连接数和连接生命周期
	d.DBBackend.DB().SetMaxIdleConns(c.DBBackend.MaxIdleConns)
	d.DBBackend.DB().SetMaxOpenConns(c.DBBackend.MaxOpenConns)
	d.DBBackend.DB().SetConnMaxLifetime(c.DBBackend.ConnMaxLifetime * time.Second)
}

// NOTE: 添加数据库，补充该方法
// @Description 关闭DAO DB创建的资源
// @Author Wangch
// @Version 1.0
// @Update Wangch 2021-09-26 init
func (d *DB) Close() {
	if d.DBBackend != nil {
		d.DBBackend.Close()
	}

	return
}

// NOTE: 添加数据库，补充该方法
// @Description 根据 model 同步指定数据库
// @Author Wangch
// @Version 1.0
// @Update Wangch 2021-09-26 init
func (d *DB) SyncDB(dbName string) {
	switch dbName {
	case model.DB_NAME_SYSTEM:
		// 业务类
		d.DBBackend.Set("gorm:table_options", "CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'NFT商品表' ROW_FORMAT = Dynamic").AutoMigrate(&model.NFTGoods{}) // NFT商品表
	default:
		fmt.Println("Unkown db name: ", dbName)
	}
	return
}
