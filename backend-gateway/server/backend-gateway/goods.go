package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/**
 * @api {get} /api/backend/goods/v1/list 获取商品列表数据
 * @apiName getGoodsListV1
 * @apiVersion 1.0.0
 * @apiGroup goods
 * @apiDescription 根据条件获取商品列表数据
 * @apiSampleRequest off
 *
 * @apiPermission 已验证用户
 *
 * @apiHeader {String} [X-Auth-Token] （预留）用户Token
 *
 * @apiParam {String} language           客户端语言 EN-英文 CN-简体中文
 * @apiParam {String} entity_id          Y星（URL参数）
 *
 * @apiParamExample {string} 请求数据格式:
 *     ?language=CN&entity_id=996
 *
 * @apiSuccess {Number} code                                错误状态码
 * @apiSuccess {String} message                             错误信息
 * @apiSuccess {Object} data                                数据详情
 * @apiSuccess {Number}   data.timestamp                    记录获取时间
 * @apiSuccess {Number}   data.total_count                  记录总条数
 * @apiSuccess {Object[]} data.tokens_have_balance                       数据详情
 * @apiSuccess {String}   data.tokens_have_balance.symbol                         币种展示名称
 *
 * @apiSuccessExample {json} 正确返回值:
 *   HTTP/1.1 200 OK
 *   {
 *       "code": 1000,
 *       "message": "",
 *       "data": {
 *           "timestamp": 1608890232,
 *           "total_count": 3,
 *   		 "tokens_have_balance": [
 *                 {
 *                     "symbol": "YTA",
 *                 }
 *   		 ]
 *       }
 *   }
 *
 * @apiUse   FailResponse
 */
func (s *RouterImpl) getGoodsListV1(c *gin.Context) {
	// 获取请求的参数
	/*
		req := &model.TokenListReq{}
		if err := c.BindQuery(req); err != nil {
			logger.Errorf(errMsg.InvalidParamsMessage, err)
			c.JSON(http.StatusPreconditionRequired, utils.ErrorResponse(utils.ErrorServiceI18nLocaleResponse(utils.InvalidParams, errMsg.InvalidParamsMessage, req.Language, err)))
			return
		}
	*/

	// 参数处理

	res, _ := s.BackendSvc.GetGoodsList()

	c.JSON(http.StatusOK, res)
	return
}
