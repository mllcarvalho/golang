package main

import "github.com/mllcarvalho/golang/go-expert-full-cycle/modulo-2/11-APIs/configs"

func main() {
	config, _ := configs.LoadConfig(".")
	println(config.DBDriver)

}
