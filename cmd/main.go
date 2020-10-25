package main

import (
  "github.com/wArrest/gin-kit/cmd/server"
  "github.com/wArrest/gin-kit/cmd/server/db"
  "github.com/wArrest/gin-kit/config"
)

func main() {
  config.Init()
  db.Init()
  server.Init()
}
