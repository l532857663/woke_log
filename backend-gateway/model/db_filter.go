package model

/* 请求数据库的参数结构 */

type GetNFTGoodsFilter struct {
	GoodsId *uint64
	Name    string
}
