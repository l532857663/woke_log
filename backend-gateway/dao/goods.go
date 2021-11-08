package dao

import "backend-gateway/model"

/* 商品相关方法 */

// @Description 获取商品列表数据
// @Author Wangch
// @Version 1.0
// @Update Wangch 2021-09-30 init
func (d *Dao) GetNFTGoods(filter *model.GetNFTGoodsFilter) ([]*model.NFTGoods, error) {
	// 查询数据库
	return d.db.GetNFTGoods(filter)
}
