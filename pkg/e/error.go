package e

import "net/http"

type Err struct {
	Status  int    `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e Err) Error() string {
	return e.Message
}

//内部错误
var InternalError = Err{
	Status:  http.StatusInternalServerError,
	Code:    40001,
	Message: "系统错误,请联系管理员",
}

//资源不存在
var DBError = Err{
	Status:  http.StatusNotFound,
	Code:    40002,
	Message: "访问的资源不存在",
}

//权限不足
var PermissionError = Err{
	Status:  http.StatusForbidden,
	Code:    40003,
	Message: "没有访问该资源的权限",
}
