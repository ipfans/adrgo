package adrgo

import (
	_ "embed"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v3"
)

var (
	bundle    *i18n.Bundle
	localizer *i18n.Localizer

	//go:embed i18n/zh-cn.yml
	zhCNBytes []byte
	//go:embed i18n/en.yml
	enBytes []byte
)

func LoadLanguage(lang string) {
	bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("yml", yaml.Unmarshal)
	bundle.MustParseMessageFileBytes(enBytes, "en.yml")
	bundle.MustParseMessageFileBytes(zhCNBytes, "zh-cn.yml")
	localizer = i18n.NewLocalizer(bundle, lang)
}

func IsValidLanguage(lang string) bool {
	tags := bundle.LanguageTags()
	for _, tag := range tags {
		if tag.String() == lang {
			return true
		}
	}
	return false
}

func T(id string) string {
	return localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID: id,
		},
	})
}
