package model

import "time"

type IDStruct struct {
	ID uint `gorm:"column:id;"`
}

// @Description 通用的基于时间的分页查询结构
// @Author Wangch
// @Version 1.0
// @Update Wangch 2021-09-26 init
type TimePageReq struct {
	StartTime int64  `form:"start_time" json:"start_time"` // 查询开始时间
	EndTime   int64  `form:"end_time" json:"end_time"`     // 查询结束时间
	PageNum   string `form:"page_num" json:"page_num"`     // 当前展示页数
	PageSize  string `form:"page_size" json:"page_size"`   // 每一页的展示数量
}

// @Description 通用的基于时间的分页查询结构
// @Author Wangch
// @Version 1.0
// @Update Wangch 2021-09-26 init
type TimePageFilter struct {
	StartTime time.Time // 查询开始时间
	EndTime   time.Time // 查询结束时间
	PageNum   int64     // 当前展示页数
	PageSize  int64     // 每一页的展示数量
	Offset    int64     // 偏移量 (page_num - 1) * page_size
}
