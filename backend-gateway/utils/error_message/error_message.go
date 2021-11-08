package errorMessage

// 错误描述信息
const (
	// 服务层
	SuccessMessage                         = "Success"                                 // 处理成功
	PongMessage                            = "Pong"                                    // pong
	InvalidParamsMessage                   = `Invalid input parameters: %v`            // 参数无效
	ReloadDbConfSuccessMessage             = "Reload dbconf success"                   // 刷新数据库配置成功
	InvalidPageNumMessage                  = `Invalid page_num, error: %v`             // 无效的 page_num
	InvalidPageSizeMessage                 = `Invalid page_size, error: %v`            // 无效的 page_size
	InvalidEntityIdMessage                 = `Invalid param entity_id: %v, error: %v`  // 无效的 entity_id
	InvalidChainIdMessage                  = `Invalid param chain_id: %v`              // 无效的 chain_id
	InvalidParamPayTxIdOrPayAddressMessage = "Invalid param pay_tx_id or pay_address"  // 付费创建参数检查
	InvalidParamRequestIdMessage           = "Invalid param request_id"                // 无效的 request_id
	InvalidParamAddressOrNameMessage       = "Invalid param account's address or name" // 查询账号地址或账号名是否可用(是否重复)
	InvalidMemoMessage                     = "Lack of memo or description"             // 缺少 memo 或 description 描述信息
	InvalidParamRecommendTypeMessage       = "Invalid param recommend"                 // 无效的 recommend
	InvalidParamTokenListMessage           = "Invalid param token symbol list"         // 无效的 token list
	InvalidParamContractOrAddressMessage   = "Invalid param contract or address"       // 缺少 contract 或 address
	InvalidParamAddressMessage             = "Invalid param address"                   // 无效的 address
	InvalidParamTokenMessage               = "Invalid param token"                     // 无效的 token
	InvalidParamPlatformIdMessage          = "Invalid param chain_id"                  // 无效的 platform_id
	InvalidParamContractMessage            = "Invalid param contract"                  // 无效的 contract
	InvalidParamPairIdMessage              = "Invalid param pair_id"                   // 无效的 pair_id

	InvalidKycStatusMessage    = "Invalid kyc status"             // 身份认证状态错误
	InvalidUrlMessage          = "Invalid url"                    // 不存在URL
	UnsupportedSetTypeMessage  = "Unsupported set type"           // 不支持的操作类型
	UnsupportedPlatformMessage = "Unsupported platform"           // 不支持的平台
	UnsupportedTokenMessage    = "Unsupported token"              // 不支持的token
	UnsupportedAddressMessage  = "Unsupported address"            // 不支持的账户地址
	UnsupportedFeeTypeMessage  = "Unsupported fee type"           // 不支持的费用类型
	UnsupportedVersionMessage  = "The current version is too low" // 当前版本过低

	SignVerificationFailedMessage = "Sign verification failed" // 验签失败
	NotHaveChainIdMessage         = "Not have chain_id"        // 未传chain_id参数

	// 数据库
	RecordNotFoundMessage    = "record not found"      // 数据库无记录
	DataAlreadyExistsMessage = "record already exists" // 主键冲突

	// 缓存
	CacheNotFoundMessage = "cache record not found or the key does not exist" // 缓存无记录或无此key
	RedisNilMessage      = "redis: nil"                                       // redis无记录，该key为nil

	// Address
	InvalidAddressMessage = "Invalid address"

	// Token
	InvalidTokenMessage  = "Invalid token"
	TokenRequiredMessage = "The request does not include the token"

	// 业务层
	InternalErrorMessage                 = "Internal error"                                      // 内部错误
	SystemErrorMessage                   = "System error, unable to obtain relevant information" // 系统错误，无法获取相关信息
	UndiscoveredRuleMessage              = "Undiscovered rule"                                   // 未发现规则设置
	DataNonexistentMessage               = "Data nonexistent"                                    // 数据不存在
	WrongRequestIdMessage                = "Wrong request id"                                    // 错误的request id
	UpdateRecordFailedMessage            = "Update record failed"                                // 更新记录失败
	PayAmountInsufficientMessage         = "Payment amount is insufficient"                      // 创建账户支付金额不够
	CheckCreateYTAAccountQuantityMessage = "Check create YTA account quantity failed"            // 已支付账户创建费用知否达标检查失败
)
