package utils

import (
	"Go-WMS/global"
	"Go-WMS/param"
	"github.com/go-playground/validator/v10"
	"strings"
)

// HandleValidatorError 处理字段校验异常
// 参数：
//		无
// 返回值：
//		c：gin.Context的指针
//		err：错误信息
func HandleValidatorError(err error) param.Resp {
	//如何返回错误信息
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		failStruct := param.Resp{
			Code: 10000,
			Msg:  global.I18nMap["10000"],
			Data: err.Error(),
		}
		return failStruct
	}
	data := removeTopStruct(errs.Translate(global.Trans))
	failStruct := param.Resp{
		Code: 10000,
		Msg:  global.I18nMap["10000"],
		Data: data,
	}
	return failStruct
}

// removeTopStruct 定义一个去掉结构体名称前缀的自定义方法
// 参数：
//		fileds：原始字段的Map
// 返回值：
//		map[string]string：处理后的Map
func removeTopStruct(fileds map[string]string) map[string]string {
	rsp := map[string]string{}
	for filed, err := range fileds {
		// 从文本的圆点(.)开始切分   处理后"mobile": "mobile为必填字段"  处理前: "PasswordLoginForm.mobile": "mobile为必填字段"
		rsp[filed[strings.Index(filed, ".")+1:]] = err
	}
	return rsp
}
