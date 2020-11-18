package utils

import "net/http"

// Response 返回数据
type Response struct {
	// 返回码
	Code int `json:"code"`
	// 返回说明
	Msg string `json:"msg"`
	// 数据
	Data interface{} `json:"data"`
}

// NewResponse 返回数据
func NewResponse(code int, msg string, data interface{}) Response {
	return Response{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

// Success 成功返回
func Success(msg string, data interface{}) Response {

	return NewResponse(http.StatusOK, msg, data)
}

// CreateFailure 新增失败
func CreateFailure(msg string, data interface{}) Response {
	return NewResponse(http.StatusAccepted, "create failure", "")
}

// BadRequest 失败返回
func BadRequest(msg string, data interface{}) Response {
	return NewResponse(http.StatusBadRequest, msg, data)
}

// Forbidden 无权限
func Forbidden(msg string, data interface{}) Response {
	return NewResponse(http.StatusForbidden, msg, data)
}

// Unauthorized 未授权
func Unauthorized(msg string, data interface{}) Response {
	return NewResponse(http.StatusUnauthorized, msg, data)
}

// ValidateFailure 校验
func ValidateFailure(data interface{}) Response {
	return NewResponse(http.StatusBadRequest, "validate failure! please see the field data for details.", data)
}
