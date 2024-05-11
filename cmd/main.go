package main

import "caipiaotong/internal/initial"

func main() {
	initial.InitConfig()
	initial.InitCache()
	initial.InitDB()
}
