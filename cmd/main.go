package main

import (
	"caipiaotong/internal/initial"
	"caipiaotong/internal/routers"
)

func main() {
	initial.InitConfig()
	initial.InitDB()
	initial.InitCache()

	r := routers.Router()
	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
