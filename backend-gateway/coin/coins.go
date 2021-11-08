package coin

import (
	"fmt"
)

// Coin 币种基础信息
type Coin struct {
	ID               uint
	ApiResource      string
	Symbol           string
	Name             string
	Decimals         uint
	BlockTime        int
	MinConfirmations int64
	SampleAddress    string
	TokenType        uint64
}

func (c *Coin) String() string {
	return fmt.Sprintf("[%s] %s (#%d)", c.Symbol, c.Name, c.ID)
}

const (
	// https://github.com/satoshilabs/slips/blob/master/slip-0044.md
	BTC = 0
	ETH = 60
	BNB = 7140
	HT  = 1010
	YTA = 1940
	EOS = 194
	OKT = 996
)

var NetWorkMap = map[string]string{
	"1":   "Ethereum Mainnet",
	"66":  "OKExChain Mainnet",
	"128": "Heco Mainnet",
	"9d7bec4bf167a7b136d0b45d8aac77bd45e761e35cbd2b7d0e88dfe05ebf3d62": "Yotta Mainnet",
	"56": "BinanceSmartChain Mainnet",

	"2":   "Morden Testnet (deprecated)",
	"3":   "Ropsten Testnet",
	"4":   "Rinkeby Testnet",
	"42":  "Kovan Testnet",
	"65":  "OKExChain Testnet",
	"256": "Heco Testnet",
	"dae29f946d62903ad5cdd6006f36f2603299a601896796f4eacd9132a6fbf949": "Yotta Testnet",
}

var CoinTypeStringMap = map[string]uint{
	"0":    Coins[BTC].ID,
	"60":   Coins[ETH].ID,
	"7140": Coins[BNB].ID,
	"1010": Coins[HT].ID,
	"1940": Coins[YTA].ID,
	"996":  Coins[OKT].ID,
}

var Coins = map[uint]Coin{
	BTC: {
		ID:               0,
		ApiResource:      "bitcoin",
		Symbol:           "BTC",
		Name:             "Bitcoin",
		Decimals:         8,
		BlockTime:        600000,
		MinConfirmations: 0,
		SampleAddress:    "bc1quvuarfksewfeuevuc6tn0kfyptgjvwsvrprk9d",
		TokenType:        0,
	},
	ETH: {
		ID:               60,
		ApiResource:      "ethereum",
		Symbol:           "ETH",
		Name:             "Ethereum",
		Decimals:         18,
		BlockTime:        10000,
		MinConfirmations: 0,
		SampleAddress:    "0xfc10cab6a50a1ab10c56983c80cc82afc6559cf1",
		TokenType:        60,
	},
	HT: {
		ID:               1010,
		ApiResource:      "heco",
		Symbol:           "HT",
		Name:             "Huobi ECO Chain",
		Decimals:         18,
		BlockTime:        3000,
		MinConfirmations: 0,
		SampleAddress:    "0xa71edc38d189767582c38a3145b5873052c3e47a",
		TokenType:        101,
	},
	YTA: {
		ID:               1940,
		ApiResource:      "yottachain",
		Symbol:           "YTA",
		Name:             "YottaCoin",
		Decimals:         4,
		BlockTime:        1,
		MinConfirmations: 0,
		SampleAddress:    "eosio.token",
		TokenType:        1,
	},
	OKT: {
		ID:               996,
		ApiResource:      "okexchain",
		Symbol:           "OKT",
		Name:             "OKExChain Token",
		Decimals:         18,
		BlockTime:        3000,
		MinConfirmations: 0,
		SampleAddress:    "",
		TokenType:        4001,
	},
	BNB: {
		ID:               7140,
		ApiResource:      "bsc",
		Symbol:           "BNB",
		Name:             "Binance Smart Chain",
		Decimals:         18,
		BlockTime:        3000,
		MinConfirmations: 0,
		SampleAddress:    "0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c",
		TokenType:        5001,
	},
	EOS: {
		ID:               194,
		ApiResource:      "eos",
		Symbol:           "EOS",
		Name:             "EOSIO",
		Decimals:         4,
		BlockTime:        1,
		MinConfirmations: 0,
		SampleAddress:    "",
		TokenType:        6001,
	},
}

func Bitcoin() Coin {
	return Coins[BTC]
}

func Ethereum() Coin {
	return Coins[ETH]
}

func Heco() Coin {
	return Coins[HT]
}

func YottaChain() Coin {
	return Coins[YTA]
}

func EOSIO() Coin {
	return Coins[EOS]
}

func OKExChain() Coin {
	return Coins[OKT]
}

func Bsc() Coin {
	return Coins[BNB]
}
