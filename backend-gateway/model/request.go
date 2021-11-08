package model

// @Description 通用参数请求结构
// @Author Wangch
// @Version 1.0
// @Update Wangch 2021-09-26 init
type CommonParamRequest struct {
	Language    string `form:"language" json:"language"`       // 客户端语言
	Application string `form:"application" json:"application"` // 应用
	Version     string `form:"version" json:"version"`         // 版本号
	Platform    string `form:"platform" json:"platform"`       // 平台
}

// @Description 通用ID请求结构
// @Author Wangch
// @Version 1.0
// @Update Wangch 2021-09-26 init
type CommonIDRequest struct {
	CommonParamRequest        // 接口请求通用参数
	EntityId           string `json:"entity_id" binding:"required"` // Entity ID
	Id                 string `json:"id"        binding:"required"` // 记录ID
	Address            string `json:"address"   binding:"required"` // 账户/地址
	ChainId            string `json:"chain_id"  binding:"required"` // 平台ID(平台主币种SLIP-0044编号)
}

// 对CommonIDRequest做验签处理的数据结构
type CommonIDData struct {
	EntityId string `json:"entity_id"` // Entity ID
	Address  string `json:"address"`   // 账户/地址
	ChainId  string `json:"chain_id"`  // 平台ID(平台主币种SLIP-0044编号)
	Id       string `json:"id"`        // 记录ID
}

// @Description 签名路由请求结构
// @Author Wangch
// @Version 1.0
// @Update Wangch 2021-09-26 init
type VerifySignatureReq struct {
	// -------------------必须参数-------------------
	EntityId string `json:"entity_id" binding:"required"` // Entity ID

	// --------------------对应参数------------------
	// 多接口共用: 创建账户、导入账户、确认账户、币种管理、账户管理
	CommonParamRequest        // 接口请求通用参数
	Address            string `json:"address,omitempty"`  // 账户/地址
	ChainId            string `json:"chain_id,omitempty"` // 平台ID(平台主币种SLIP-0044编号)
}
