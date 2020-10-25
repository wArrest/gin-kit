package controller

import (
  "github.com/gin-gonic/gin"
)

type UserController struct {
}
//获取用户名
func (u UserController) GetUserName(c *gin.Context) {
  return
}

//用户注册
func (u UserController) SignUp(c *gin.Context) {
  //todo 处理和验证参数
  //us:=service.UserService{}
  //us.SignUp(c,nil)
  return
}
