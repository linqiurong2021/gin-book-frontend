package validator

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/linqiurong2021/gin-book-frontend/utils"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

// Trans 定义一个全局翻译器T
var Trans ut.Translator

// InitTrans 初始化翻译器
func InitTrans(locale string) (err error) {
	// 修改gin框架中的Validator引擎属性，实现自定制
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {

		zhT := zh.New() // 中文翻译器
		enT := en.New() // 英文翻译器

		// 第一个参数是备用（fallback）的语言环境
		// 后面的参数是应该支持的语言环境（支持多个）
		// uni := ut.New(zhT, zhT) 也是可以的
		uni := ut.New(enT, zhT, enT)

		// locale 通常取决于 http 请求头的 'Accept-Language'
		var ok bool
		// 也可以使用 uni.FindTranslator(...) 传入多个locale进行查找
		Trans, ok = uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s) failed", locale)
		}

		//注册一个函数，获取struct tag里自定义的label作为字段名
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := fld.Tag.Get("label")
			return name
		})

		// 注册翻译器
		switch locale {
		case "en":
			err = enTranslations.RegisterDefaultTranslations(v, Trans)
		case "zh":
			err = zhTranslations.RegisterDefaultTranslations(v, Trans)
		default:
			err = enTranslations.RegisterDefaultTranslations(v, Trans)
		}
		return
	}
	return
}

// Translate 转义
func Translate(errs validator.ValidationErrors) validator.ValidationErrorsTranslations {
	return errs.Translate(Trans)
}

// ValidateError 校验是否是验证的错误
func ValidateError(err error) (errs validator.ValidationErrors, ok bool) {
	errs, ok = err.(validator.ValidationErrors)
	return
}

// Validate 校验登录表单
func Validate(c *gin.Context, err error) bool {
	//
	// 参数校验判断
	if err != nil {
		errs, ok := ValidateError(err)
		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			c.JSON(http.StatusBadRequest, utils.BadRequest(err.Error(), ""))
			return false
		}
		// validator.ValidationErrors类型错误则进行翻译
		c.JSON(http.StatusBadGateway, utils.ValidateFailure(Translate(errs)))
		return false
	}
	return true
}
