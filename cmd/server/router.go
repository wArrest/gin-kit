package server

import (
  "github.com/gin-gonic/gin"
  "github.com/wArrest/gin-kit/controller"
)

func NewRouter() *gin.Engine {
  router := gin.New()
  router.Use(gin.Logger())
  router.Use(gin.Recovery())

  v1 := router.Group("v1")
  {
    userGroup := v1.Group("user")
    {
      user := new(controller.UserController)
      userGroup.GET("/:name", user.GetUserName)
    }

  }
  return router
}
