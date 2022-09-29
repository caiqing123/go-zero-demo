package validator

import (
	"strings"

	en_lang "github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh_Hans"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

type Validator struct {
	Validator *validator.Validate
	Uni       *ut.UniversalTranslator
	Trans     map[string]ut.Translator
}

func NewValidator() *Validator {
	v := Validator{}
	en := en_lang.New()
	zh := zh_Hans.New()
	v.Uni = ut.New(zh, en, zh)
	v.Validator = validator.New()
	enTrans, _ := v.Uni.GetTranslator("en")
	zhTrans, _ := v.Uni.GetTranslator("zh")
	v.Trans = make(map[string]ut.Translator)
	v.Trans["en"] = enTrans
	v.Trans["zh"] = zhTrans

	err := en_translations.RegisterDefaultTranslations(v.Validator, enTrans)
	if err != nil {
		return nil
	}
	err = zh_translations.RegisterDefaultTranslations(v.Validator, zhTrans)
	if err != nil {
		return nil
	}

	return &v
}

func (v *Validator) Validate(data interface{}, lang string) string {
	err := v.Validator.Struct(data)
	if err == nil {
		return ""
	}

	errs, ok := err.(validator.ValidationErrors)
	if ok {
		transData := errs.Translate(v.Trans[lang])
		s := strings.Builder{}
		for _, v := range transData {
			s.WriteString(v)
			s.WriteString(" ")
		}
		return s.String()
	}

	invalid, ok := err.(*validator.InvalidValidationError)
	if ok {
		return invalid.Error()
	}

	return ""
}
