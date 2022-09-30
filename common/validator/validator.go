package validator

import (
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/Masterminds/squirrel"
	en_lang "github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh_Hans"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
)

type Validator struct {
	Validator *validator.Validate
	Uni       *ut.UniversalTranslator
	Trans     map[string]ut.Translator
}

var DBModel sqlc.CachedConn

var translations = []struct {
	tag             string
	translation     string
	override        bool
	customRegisFunc validator.RegisterTranslationsFunc
	customTransFunc validator.TranslationFunc
}{
	{
		tag:         "exists",
		translation: "{0}不存在",
		override:    false,
	},
	{
		tag:         "not_exists",
		translation: "{0}已存在",
		override:    false,
	},
}

func NewValidator() *Validator {
	v := Validator{}
	en := en_lang.New()
	zh := zh_Hans.New()
	v.Uni = ut.New(zh, en, zh)
	v.Validator = validator.New()
	//注册自定义的校验函数 exists
	err := v.Validator.RegisterValidation("exists", exists)
	if err != nil {
		return nil
	}
	err = v.Validator.RegisterValidation("not_exists", not_exists)
	if err != nil {
		return nil
	}

	// 注册 RegisterTagNameFunc 根据label字段名称处理
	v.Validator.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := strings.SplitN(field.Tag.Get("label"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	enTrans, _ := v.Uni.GetTranslator("en")
	zhTrans, _ := v.Uni.GetTranslator("zh")

	//自定义方法翻译
	for _, t := range translations {

		if t.customTransFunc != nil && t.customRegisFunc != nil {

			err = v.Validator.RegisterTranslation(t.tag, zhTrans, t.customRegisFunc, t.customTransFunc)

		} else if t.customTransFunc != nil && t.customRegisFunc == nil {

			err = v.Validator.RegisterTranslation(t.tag, zhTrans, registrationFunc(t.tag, t.translation, t.override), t.customTransFunc)

		} else if t.customTransFunc == nil && t.customRegisFunc != nil {

			err = v.Validator.RegisterTranslation(t.tag, zhTrans, t.customRegisFunc, translateFunc)

		} else {
			err = v.Validator.RegisterTranslation(t.tag, zhTrans, registrationFunc(t.tag, t.translation, t.override), translateFunc)
		}

		if err != nil {
			return nil
		}
	}

	v.Trans = make(map[string]ut.Translator)
	v.Trans["en"] = enTrans
	v.Trans["zh"] = zhTrans

	err = en_translations.RegisterDefaultTranslations(v.Validator, enTrans)
	if err != nil {
		return nil
	}
	err = zh_translations.RegisterDefaultTranslations(v.Validator, zhTrans)
	if err != nil {
		return nil
	}

	return &v
}

func registrationFunc(tag string, translation string, override bool) validator.RegisterTranslationsFunc {

	return func(ut ut.Translator) (err error) {

		if err = ut.Add(tag, translation, override); err != nil {
			return
		}

		return

	}

}

func translateFunc(ut ut.Translator, fe validator.FieldError) string {

	t, err := ut.T(fe.Tag(), fe.Field())
	if err != nil {
		log.Printf("警告: 翻译字段错误: %#v", fe)
		return fe.(error).Error()
	}

	return t
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

func getValue(pt reflect.Value) interface{} {
	var value interface{}
	//根据类型取值
	switch pt.Kind() {
	case reflect.String, reflect.Array:
		value = pt.String()
	case reflect.Bool:
		value = pt.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		value = pt.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		value = pt.Uint()
	case reflect.Float32, reflect.Float64:
		value = pt.Float()
	}

	return value
}

// 自定义的校验规则
func exists(f validator.FieldLevel) bool {
	value := getValue(f.Field())

	//查询处理
	query, values, err := squirrel.Select("COUNT(*)").From(f.Param()).Where(squirrel.Eq{f.StructFieldName(): fmt.Sprintf("%v", value)}).ToSql()
	if err != nil {
		return false
	}
	var resp int64
	err = DBModel.QueryRowNoCache(&resp, query, values...)
	if err != nil || resp == 0 {
		return false
	}
	return true
}

func not_exists(f validator.FieldLevel) bool {
	value := getValue(f.Field())

	rng := strings.Split(f.Param(), "-")
	table := rng[0]

	sql := squirrel.Select("COUNT(*)").From(table).Where(squirrel.Eq{f.StructFieldName(): fmt.Sprintf("%v", value)})

	if len(rng) > 1 {
		filedID := rng[1]
		exceptID := f.Parent().FieldByName(rng[1])
		value1 := getValue(exceptID)
		sql = sql.Where(filedID+" != ?", value1)
	}

	//查询处理
	query, values, err := sql.ToSql()
	if err != nil {
		return false
	}
	var resp int64
	err = DBModel.QueryRowNoCache(&resp, query, values...)

	if err != nil || resp != 0 {
		return false
	}
	return true
}
