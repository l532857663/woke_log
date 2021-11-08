package db

import (
	"backend-gateway/model"
	"backend-gateway/utils"
)

// @Description 获取商品列表数
// @Author Wangch
// @Version 1.0
// @Update Wangch 2021-09-30 init
func (db *DB) GetNFTGoods(filter *model.GetNFTGoodsFilter) ([]*model.NFTGoods, error) {
	var list []*model.NFTGoods
	err := db.DBBackend.Model(&model.NFTGoods{}).
		Where(&model.NFTGoods{
			GoodsId: filter.GoodsId,
			Name:    filter.Name,
		}).Find(&list).Error
	if err != nil {
		return list, utils.DbErrorTransform(err)
	}
	return list, nil
}
