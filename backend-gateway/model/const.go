package model

const (
	/*************************************** 系统常量 **********************************/
	// AppletName
	APPLET_NAME_BACKEND = "backend-gateway"
	// ServerName
	SERVER_NAME_BACKEND_GATEWAY = "backend-gateway"

	// Panic stack size
	PANIC_STACK_SIZE = 4096

	// Language
	LANGUAGE_KEY_NAME           = "language"
	LANGUAGE_DEFAULT            = LANGUAGE_ENGLISH
	LANGUAGE_ENGLISH            = "EN"
	LANGUAGE_SIMPLIFIED_CHINESE = "CN"

	// Cache
	CACHE_NAME_REDIS = "Redis"

	// DbName
	DB_NAME_SYSTEM = "backend_system"

	// 时间分页参数flag
	REQ_FLAG_ALL  = "all"
	REQ_FLAG_TIME = "time"
	REQ_FLAG_PAGE = "page"

	// GetChainTokenCheckCache
	CHECK_CACHE_FALSE = "0" // 不查缓存
	CHECK_CACHE_TRUE  = "1" // 查缓存

	// 货币代码
	CURRENCY_CODE_CNY = "CNY" // 人民币
	CURRENCY_CODE_USD = "USD" // 美元

	/*************************************** 常用参数 **********************************/
	// IsDefaultView
	IS_DEFAULT_VIEW_TRUE = "1"

	// AvailableState
	AVAILABLE_STATE_FALSE = "0" // 不可用
	AVAILABLE_STATE_TRUE  = "1" // 可用

	// ActiveState
	ACTIVE_STATE_FALSE = "0" // 不可用
	ACTIVE_STATE_TRUE  = "1" // 可用

	// SupportState
	SUPPORT_STATE_FALSE = "0" // 不支持
	SUPPORT_STATE_TRUE  = "1" // 支持

	// PermissionName
	PERMISSION_NAME_ACTIVE = "active"
	PERMISSION_NAME_OWNER  = "owner"

	// KycType
	KYC_TYPE_SEMI_ANONYMOUS string = "0" // 半匿名
	KYC_TYPE_ANONYMOUS      string = "1" // 匿名
	KYC_TYPE_SEMI_REAL      string = "3" // 半实名
	KYC_TYPE_REAL           string = "4" // 实名
)
