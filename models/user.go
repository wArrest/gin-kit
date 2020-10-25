package models

//用户注册表单
type UserSignUp struct {
  Name  string `json:"name"`
  Email string `json:"email"`
}

//用户信息
type UserInfo struct {
  UUID      string `json:"uuid"`
  Name      string `json:"name"`
  Email     string `json:"email"`
  CreatedAt int    `json:"created_at"`
}
