package main

import (
	"caipiaotong/configs"
	"caipiaotong/internal/cache"
	"caipiaotong/internal/dao"
	"caipiaotong/internal/routers"
)

func main() {
	configs.InitConfig()
	cache.InitCache()
	dao.InitDB()
	r := routers.NewRouter()
	r.Run(":8080")
}
