package main

import (
	"caipiaotong/configs"
	"caipiaotong/internal/cache"
	"caipiaotong/internal/dao"
	"caipiaotong/internal/routers"
)

func main() {

	r := routers.Router()
	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
func init() {
	configs.InitConfig()
	cache.InitCache()
	dao.InitDB()
}
