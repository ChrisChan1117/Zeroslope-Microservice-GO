package main

import (
	"fmt"

	"./router"
	"./utilities"
)

// @title ZeroSlope API
// @version 1.0
// @description Zeroslope microservice architecture written in go
func main() {
	fmt.Println(">> Loading configuration ...")
	cfg := utilities.LoadConfigFromPath("src/config.yaml")

	fmt.Println(">> Loading routes ...")
	router := router.SetupRouter(cfg)

	fmt.Println(">> Starting service ...")
	router.Run(cfg.Server.Port)
}
