package config

import "time"

//	local:		本地环境
//	joindebug:	联调环境
//	dev:		内部测试环境
//	beta:		外部测试环境(灰度环境)
//	prod:		生产环境
const (
	RUN_ENV_LOCAL      string = "local"
	RUN_ENV_JOIN_DEBUG string = "joindebug"
	RUN_ENV_DEV        string = "dev"
	RUN_ENV_BETA       string = "beta"
	RUN_ENV_PROD       string = "prod"
)

// 服务全局配置
type ServerGeneralConfig struct {
	Name          string     // 服务名称
	RunEnv        string     // 服务运行模式
	ListenAddress string     // 服务监听地址
	ListenPort    uint16     // 服务监听端口
	PProfConf     *PProfConf // PProf服务配置
}

// PProf 分析工具专用配置
type PProfConf struct {
	Enabled bool
	Address string // Pprof服务地址
	Port    uint16 // Pprof服务端口
}

// 数据库配置
type DatabaseConfig struct {
	Driver          string        // 数据库驱动：mysql or postgre
	Address         string        // 数据库地址：ip:port
	DBname          string        // 数据库名称
	User            string        // 数据库用户名
	Password        string        // 数据库密码
	Charset         string        // 数据库字符集
	ParseTime       string        // 解析时间类型
	Dsn             string        // 数据源名称 Data Source Name: network:[//[username[:password]@]address[:port][,address[:port]]][/path][?query][#fragment]
	MaxIdleConns    int           // 最大空闲连接数
	MaxOpenConns    int           // 最大打开连接数 0为不限制
	ConnMaxLifetime time.Duration // 连接最大生命周期 0为永远不断开
}

// 缓存配置
type CacheConfig struct {
	Enabled bool
	Engine  string
	Redis   *RedisConfig
}

// Redis 配置
type RedisConfig struct {
	Enable       bool
	Cluster      bool
	Addresses    []string
	AddressIndex int
	Credential   string
	Db           int
	PoolSize     int
	PoolTimeout  string
	IdleTimeout  string
}

// blockchain配置
type BlockChainConfig struct {
	BlockbookApiUrl           string        // Blockbook 服务rest api url
	ApiRpcUrl                 string        // api rpc 节点地址
	ApiKey                    string        // api rpc 节点key
	IsBasicAuth               bool          // rpc节点是否启用 basic auth 验证方式
	RpcUser                   string        // rpc 访问用户名
	RpcPassword               string        // rpc 访问密码
	GetGasTimeOut             time.Duration // 获取gas费超时时间（秒）
	GetTokenBalanceExpiration time.Duration // 获取token余额过期时间（秒）
	EthDefaultGasLimit        uint64
	EthGasTrackerURL          string // 获取gas费接口URL
	EthGasTrackerURLBackup    string // 获取gas费接口URL
	EthGasTrackerApiKey       string // 获取gas费接口api key
	EthGasTrackerAppName      string // 获取gas费接口app name
	IconUrlTemplate           string // 获取icon_url的模板链接
}
