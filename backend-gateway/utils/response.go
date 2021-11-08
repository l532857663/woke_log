package utils

import (
	errMsg "backend-gateway/utils/error_message"
)

// 服务层通用应答结构
type CommonResponse struct {
	Code    int32       `json:"code,omitempty"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// 服务层生成成功应答
func SuccessResponse(data interface{}) *CommonResponse {
	var res = CommonResponse{}
	res.Code = SuccessCode
	res.Message = errMsg.SuccessMessage
	res.Data = data
	return &res
}

// 服务层生成错误应答
func ErrorResponse(code int32, msg string) *CommonResponse {
	var res = CommonResponse{}
	res.Code = code
	res.Message = formatErrorMsg(msg)
	return &res
}
