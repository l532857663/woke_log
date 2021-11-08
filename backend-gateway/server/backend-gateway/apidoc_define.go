package main

// 本文件仅用于定义 apidoc 中的 apiDefine 请勿用于其他用途!

/**
 * 统一应答基础数据
 * @apiDefine CommonResponseData
 *
 * @apiSuccess {Number} code      错误状态码
 * @apiSuccess {String} message   错误信息
 * @apiSuccess {Object} data      数据详情
 */

/**
 * 统一正确应答（空数据）
 * @apiDefine CommonSuccessResponse
 *
 * @apiSuccessExample {json} 正确返回值:
 *   HTTP/1.1 200 OK
 *   {
 *        "code": 1000,
 *        "message": "1",
 *        "data": {}
 *   }
 */

/**
 * 统一错误应答
 * @apiDefine FailResponse
 *
 * @apiError InvalidParams_1300       无效参数
 * @apiError RecordNotFound_1301      记录不存在
 * @apiError DataAlreadyExists_1302   数据重复
 * @apiError InvalidKycStatus_1330    身份认证状态错误
 * @apiError InvalidInviteCode_1331   无效的邀请码
 * @apiError InvalidInviteAction_1332 禁止自己邀请自己
 * @apiError DataAlreadyExists_1399   服务内部错误
 *
 * @apiErrorExample {json} 错误返回值:
 *   HTTP/1.1 500 Internal Server Error
 *   {
 *        "code": 1301,
 *        "message": "record not found"
 *   }
 */
