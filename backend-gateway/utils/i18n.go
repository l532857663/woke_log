package utils

import (
	"backend-gateway/model"
	_ "backend-gateway/utils/i18n/catalog"

	"github.com/mylukin/easy-i18n/i18n"
	"golang.org/x/text/language"
)

var (
	pEN      *i18n.Printer
	pCN      *i18n.Printer
	pDefault *i18n.Printer
)

func init() {
	pEN = i18n.NewPrinter(language.English)
	pCN = i18n.NewPrinter(language.SimplifiedChinese)

	pDefault = pEN
}

// 按语言输出最终应答文本
func GetI18nLocaleMessage(message, lang string, args ...interface{}) string {
	var msg string

	switch lang {
	case model.LANGUAGE_ENGLISH:
		msg = pEN.Sprintf(message, args...)
	case model.LANGUAGE_SIMPLIFIED_CHINESE:
		msg = pCN.Sprintf(message, args...)
	default:
		msg = pDefault.Sprintf(message, args...)
	}

	return msg
}
