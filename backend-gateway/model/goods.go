package model

import "github.com/jinzhu/gorm"

/*  商品表 */

type NFTGoods struct {
	gorm.Model
	GoodsId *uint64 `json:"-" gorm:"column:goods_id; type:int(10);      not null; default:1  comment '商品ID';   unique_index: uidx_goods_id;"` // 商品ID
	Name    string  `json:"-" gorm:"column:name;     type:varchar(50);  not null; default:'' comment '商品名称';"`                                // 商品名称
	Price   float64 `json:"-" gorm:"column:price;    type:float;        not null; default:0  comment '商品价格';"`                                // 商品价格
	ImgUrl  string  `json:"-" gorm:"column:img_url;  type:varchar(200); not null; default:'' comment '商品图片链接';"`                              // 商品图片链接
}
