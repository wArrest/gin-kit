package server

import "github.com/wArrest/gin-kit/config"

func Init() {
	cfg := config.GetConfig()

	r := NewRouter()
	r.Run(cfg.Port)

}
