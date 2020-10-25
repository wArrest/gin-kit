package service

import (
  "context"
  "github.com/friendsofgo/errors"
  "github.com/volatiletech/sqlboiler/v4/boil"
  "github.com/volatiletech/sqlboiler/v4/queries/qm"
  "github.com/wArrest/gin-kit/cmd/server/db"
  "github.com/wArrest/gin-kit/entity"
  "github.com/wArrest/gin-kit/models"
)

type UserService struct {
}


//用户注册服务
func (us UserService) SignUp(ctx context.Context, userSignUp *models.UserSignUp) (*models.UserInfo, error) {

  if exist, _ := entity.Users(qm.Where("email=?", userSignUp.Email)).Exists(ctx, db.GetDB()); exist {
    return nil, errors.New("该用户已存在")
  }
  user := entity.User{
    Name:  userSignUp.Name,
    Email: userSignUp.Email,
  }
  if err := user.Insert(ctx, db.GetDB(), boil.Infer()); err != nil {
    return nil, err
  }

  var userInfo *models.UserInfo
  if err := entity.Users(qm.Where("id=?", user.ID)).Bind(ctx, db.GetDB(), userInfo); err != nil {
    return nil, err
  }

  return userInfo, nil
}
