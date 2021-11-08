package dao

import (
	"backend-gateway/model"

	"github.com/jinzhu/gorm"
)

// 数据处理
type Storage interface {
	// Control
	GetDbTxn(dbName string) *gorm.DB
	GetDbConn(dbName string) *gorm.DB
	SyncDB(dbName string)
	Close()

	// system
	GetNFTGoods(filter *model.GetNFTGoodsFilter) ([]*model.NFTGoods, error)
}
