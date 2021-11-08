package utils

import (
	"errors"
	"fmt"
	"strings"

	errMsg "backend-gateway/utils/error_message"

	"github.com/go-redis/redis/v7"
)

// 错误码
const (
	// 业务服务
	SuccessCode      = 1000 // 处理成功
	InvalidParams    = 1300 // 无效参数
	InvalidKycStatus = 1330 // 身份认证状态错误
	InternalError    = 1999 // 内部错误

	// 数据库
	RecordNotFound    = 1301 // 数据不存在，无数据
	DataAlreadyExists = 1302 // 数据已经存在，主键重复
)

var (
	// ErrTokenRequired 请求中未包含Token
	ErrTokenRequired = errors.New(errMsg.TokenRequiredMessage)
	// ErrInvalidToken Token验证失败
	ErrInvalidToken = errors.New(errMsg.InvalidTokenMessage)
	// ErrInvalidKycStatus 身份认证状态不符合
	ErrInvalidKycStatus = errors.New(errMsg.InvalidKycStatusMessage)
	// ErrInvalidUrl 不存在URL
	ErrInvalidUrl = errors.New(errMsg.InvalidUrlMessage)
	// ErrUnsupportedSetType 不支持的操作类型
	ErrUnsupportedSetType = errors.New(errMsg.UnsupportedSetTypeMessage)
	// ErrUnsupportedPlatform 不支持的平台
	ErrUnsupportedPlatform = errors.New(errMsg.UnsupportedPlatformMessage)
	// ErrUnsupportedToken 不支持的token
	ErrUnsupportedToken = errors.New(errMsg.UnsupportedTokenMessage)
	// ErrUnsupportedAddress 不支持的账户地址
	ErrUnsupportedAddress = errors.New(errMsg.UnsupportedAddressMessage)
	// ErrUnsupportedFeeType 不支持的费用类型
	ErrUnsupportedFeeType = errors.New(errMsg.UnsupportedFeeTypeMessage)
	// ErrDataNonexistent 数据不存在
	ErrDataNonexistent = errors.New(errMsg.DataNonexistentMessage)
	// ErrWrongRequestId 错误的请求id（该Key非当前request id对应的goroutine所设置）
	ErrWrongRequestId = errors.New(errMsg.WrongRequestIdMessage)
	// ErrUpdateRecord 更新记录失败
	ErrUpdateRecordFailed = errors.New(errMsg.UpdateRecordFailedMessage)
	// ErrPayAmountInsufficient 创建账户支付金额不够
	ErrPayAmountInsufficient = errors.New(errMsg.PayAmountInsufficientMessage)
	// ErrCheckCreateYTAAccountQuantity 已支付账户创建费用知否达标检查失败
	ErrCheckCreateYTAAccountQuantity = errors.New(errMsg.CheckCreateYTAAccountQuantityMessage)
	// ErrJsonMarshalError
	ErrJsonMarshalError = errors.New("Json marshal error")
	// ErrJsonUnMarshalError
	ErrJsonUnMarshalError = errors.New("Json unmarshal error")

	// TypePlatformUnmarshal
	ErrTypePlatformUnmarshal = errors.New("Platform unmarshal error")
	// TypePlatformNormalize
	ErrTypePlatformNormalize = errors.New("Platform normalize error")
	// TypePlatformUnknown
	ErrTypePlatformUnknown = errors.New("Platform unknown error")
	// TypePlatformClient
	ErrTypePlatformClient = errors.New("Platform client generic error")
	// TypePlatformError
	ErrTypePlatformError = errors.New("Custom platform error")
	// TypePlatformApi
	ErrTypePlatformApi = errors.New("Platform API error")
	// TypeUnknown
	ErrTypeUnknown = errors.New("Unknown error")

	// ErrUnregisteredName
	ErrUnregisteredName = errors.New("Unregistered name or resolver not set")
	// ErrInvalidAddress
	ErrInvalidAddress = errors.New(errMsg.InvalidAddressMessage)
	// ErrInvalidAddressLength
	ErrInvalidAddressLength = errors.New("Invalid address length")
	// ErrInvalidResultLength
	ErrInvalidResultLength = errors.New("Invalid result length")
)

// TypePlatformRequest:
func ErrTypePlatformRequest(err error) error {
	return errors.New(fmt.Sprint("Platform request error: ", err))
}

// ErrRPCGetObjectMarshalError
func ErrRPCGetObjectMarshalError(err error) error {
	return errors.New(fmt.Sprint("Json-rpc get object marshal error: ", err))
}

// ErrRPCGetObjectUnMarshalError
func ErrRPCGetObjectUnMarshalError(err error) error {
	return errors.New(fmt.Sprint("Json-rpc get object unmarshal error: ", err))
}

var StrToCode = map[string]int32{
	errMsg.RecordNotFoundMessage:    RecordNotFound,
	errMsg.DataAlreadyExistsMessage: DataAlreadyExists,
	errMsg.InternalErrorMessage:     InternalError,
}

// NOTE:
// 第三方包和一些常规的错误封装
// 请不要直接使用第三方包的报错信息
// 第三方包升级后，只需修改对应此处的封装处理，不会影响到业务层
const (
	// 数据库
	// 主键冲突
	MysqlErrorDuplicateKey = "1062" // 主键（唯一索引）冲突
)

var (
	// 业务服务

	// 数据库
	ErrRecordNotFound    = errors.New(errMsg.RecordNotFoundMessage) // gorm.ErrRecordNotFound
	ErrDataAlreadyExists = errors.New(errMsg.DataAlreadyExistsMessage)

	// 缓存
	RedisNil         = redis.Nil
	ErrCacheNotFound = errors.New(errMsg.CacheNotFoundMessage)
)

func GenerateError(msg string) error {
	return errors.New(msg)
}

func formatErrorMsg(msg string) string {
	var errMsg string
	if msg != "" {
		errMsg = fmt.Sprintf("%s", msg)
	}

	return errMsg
}

// 数据库错误转换
func DbErrorTransform(err error) error {
	if err == nil {
		return nil
	}

	if strings.Contains(err.Error(), errMsg.RecordNotFoundMessage) {
		return ErrRecordNotFound
	} else if strings.Contains(err.Error(), MysqlErrorDuplicateKey) {
		return ErrDataAlreadyExists
	}

	// 如果找不到匹配项返回原错误
	return err
}

// 缓存错误转换
func CacheErrorTransform(err error, key string) error {
	if err == nil {
		return nil
	}

	if strings.Contains(err.Error(), RedisNil.Error()) {
		return errors.New(fmt.Sprintf("%s, key: %s", errMsg.CacheNotFoundMessage, key))
	}

	// 如果找不到匹配项返回原错误
	return err
}

// @Description 业务应答数据（多国语言 i18n locale 支持）
// @Author Ryen
// @Version 1.0
// @Update Ryen 2020-11-14 init
func ErrorServiceI18nLocaleResponse(code int32, message, lang string, args ...interface{}) (int32, string) {
	var localeMessage string
	localeMessage = GetI18nLocaleMessage(message, lang, args...)

	return code, localeMessage
}

func ErrorTypeJudgment(err error) (int32, string) {
	if strings.Contains(err.Error(), errMsg.RecordNotFoundMessage) {
		return StrToCode[errMsg.RecordNotFoundMessage], errMsg.RecordNotFoundMessage
	} else if strings.Contains(err.Error(), errMsg.DataAlreadyExistsMessage) {
		return StrToCode[errMsg.DataAlreadyExistsMessage], errMsg.DataAlreadyExistsMessage
	}

	return StrToCode[errMsg.InternalErrorMessage], fmt.Sprintf("%s: %+v", errMsg.InternalErrorMessage, err)
}

// ErrRPCCallError
func ErrRPCCallError(code int, message string) error {
	errStr := fmt.Sprintf("RPC call error, code: %d, message: %s", code, message)
	return errors.New(errStr)
}
