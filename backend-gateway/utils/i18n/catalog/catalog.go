package catalog

import (
	errMsg "backend-gateway/utils/error_message"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// init
func init() {
	initEn(language.Make("en"))
	initZhHans(language.Make("zh-Hans"))
}

// initEn will init en support.
func initEn(tag language.Tag) {
	// Server 层
	message.SetString(tag, errMsg.PongMessage, errMsg.PongMessage)                                                       // ping pong
	message.SetString(tag, errMsg.ReloadDbConfSuccessMessage, errMsg.ReloadDbConfSuccessMessage)                         // Reload dbconf success
	message.SetString(tag, errMsg.InvalidParamsMessage, errMsg.InvalidParamsMessage)                                     // Invalid input parameters: %v
	message.SetString(tag, errMsg.InvalidPageNumMessage, errMsg.InvalidPageNumMessage)                                   // Invalid page_num, error: %v
	message.SetString(tag, errMsg.InvalidPageSizeMessage, errMsg.InvalidPageSizeMessage)                                 // Invalid page_size, error: %v
	message.SetString(tag, errMsg.InvalidEntityIdMessage, errMsg.InvalidEntityIdMessage)                                 // Invalid param entity_id: %v
	message.SetString(tag, errMsg.InvalidParamPayTxIdOrPayAddressMessage, errMsg.InvalidParamPayTxIdOrPayAddressMessage) // Invalid param pay_tx_id or pay_address
	message.SetString(tag, errMsg.InvalidParamRequestIdMessage, errMsg.InvalidParamRequestIdMessage)                     // Invalid param request_id
	message.SetString(tag, errMsg.InvalidParamAddressOrNameMessage, errMsg.InvalidParamAddressOrNameMessage)             // Invalid param account's address or name
	message.SetString(tag, errMsg.InvalidMemoMessage, errMsg.InvalidMemoMessage)                                         // Lack of memo or description
	message.SetString(tag, errMsg.InvalidParamRecommendTypeMessage, errMsg.InvalidParamRecommendTypeMessage)             // Invalid param recommend
	message.SetString(tag, errMsg.InvalidParamTokenListMessage, errMsg.InvalidParamTokenListMessage)                     // Invalid param token symbol list
	message.SetString(tag, errMsg.InvalidParamContractOrAddressMessage, errMsg.InvalidParamContractOrAddressMessage)     // Invalid param contract or address
	message.SetString(tag, errMsg.InvalidParamAddressMessage, errMsg.InvalidParamAddressMessage)                         // Invalid param address
	message.SetString(tag, errMsg.InvalidParamTokenMessage, errMsg.InvalidParamTokenMessage)                             // Invalid param token
	message.SetString(tag, errMsg.InvalidParamPlatformIdMessage, errMsg.InvalidParamPlatformIdMessage)                   // Invalid param platform_id
	message.SetString(tag, errMsg.InvalidParamContractMessage, errMsg.InvalidParamContractMessage)                       // Invalid param contract
	message.SetString(tag, errMsg.InvalidParamPairIdMessage, errMsg.InvalidParamPairIdMessage)                           // Invalid param pair_id
	message.SetString(tag, errMsg.SignVerificationFailedMessage, errMsg.SignVerificationFailedMessage)                   // Sign verification failed

	// Service 层
}

// initZhHans will init zh-Hans support.
func initZhHans(tag language.Tag) {
	// Server 层
	message.SetString(tag, errMsg.PongMessage, "链接正常")                                                   // ping pong
	message.SetString(tag, errMsg.ReloadDbConfSuccessMessage, "成功刷新数据库配置")                               // Reload dbconf success
	message.SetString(tag, errMsg.InvalidParamsMessage, "请求参数错误: %v")                                    // Invalid input parameters: %v
	message.SetString(tag, errMsg.InvalidPageNumMessage, "请求参数page_num无效, 错误信息: %v")                     // Invalid page_num, error: %v
	message.SetString(tag, errMsg.InvalidPageSizeMessage, "请求参数page_size无效, 错误信息: %v")                   // Invalid page_size, error: %v
	message.SetString(tag, errMsg.InvalidEntityIdMessage, "请求参数entity_id: %v无效, 错误信息: %v")               // Invalid param entity_id: %v
	message.SetString(tag, errMsg.InvalidParamPayTxIdOrPayAddressMessage, "请求参数pay_tx_id或pay_address无效") // Invalid param pay_tx_id or pay_address
	message.SetString(tag, errMsg.InvalidParamRequestIdMessage, "请求参数request_id无效")                      // Invalid param request_id
	message.SetString(tag, errMsg.InvalidParamAddressOrNameMessage, "缺少address或name参数")                  // Invalid param account's address or name
	message.SetString(tag, errMsg.InvalidMemoMessage, "缺少memo或description参数")                            // Lack of memo or description
	message.SetString(tag, errMsg.InvalidParamRecommendTypeMessage, "请求参数recommend无效")                   // Invalid param recommend
	message.SetString(tag, errMsg.InvalidParamTokenListMessage, "请求参数token_symbol无效")                    // Invalid param token symbol list
	message.SetString(tag, errMsg.InvalidParamContractOrAddressMessage, "缺少contract或address参数")          // Invalid param contract or address
	message.SetString(tag, errMsg.InvalidParamAddressMessage, "请求参数address无效")                           // Invalid param address
	message.SetString(tag, errMsg.InvalidParamTokenMessage, "请求参数token无效")                               // Invalid param token
	message.SetString(tag, errMsg.InvalidParamPlatformIdMessage, "请求参数chain_id无效")                       // Invalid param platform_id
	message.SetString(tag, errMsg.InvalidParamContractMessage, "请求参数contract无效")                         // Invalid param contract
	message.SetString(tag, errMsg.InvalidParamPairIdMessage, "请求参数pair_id无效")                            // Invalid param pair_id
	message.SetString(tag, errMsg.SignVerificationFailedMessage, "验签失败，请检查请求参数")                         // Sign verification failed

	// Service 层
}
