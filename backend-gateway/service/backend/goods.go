package backend

import (
	"backend-gateway/library/config"
	"backend-gateway/model"
	"backend-gateway/utils"
	"fmt"
)

// @Description 获取商品列表
// @Author Wangch
// @Version 1.0
// @Update Wangch 2021-10-08 init
func (s *BackendService) GetGoodsList() (*utils.CommonResponse, error) {
	resp := &utils.CommonResponse{}

	// 设置 mock dao 的期望值
	if s.Conf.ServerGeneral.RunEnv == config.RUN_ENV_JOIN_DEBUG {
	}

	fmt.Printf("wch----- 测试数据")

	respData := &model.GetGoodslistResp{}

	resp = utils.SuccessResponse(respData)
	return resp, nil
}
