package main

import (
	"fmt"

	"github.com/cfryerdev/zeroslope-golang/pkg/config"
	"github.com/cfryerdev/zeroslope-golang/pkg/database"
)

// @title ZeroSlope API
// @version 1.0
// @description Zeroslope microservice architecture written in go
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	fmt.Println(">> Loading configuration ...")
	cfg := config.LoadConfigFromPath("config.yaml")

	fmt.Println(">> Loading routes ...")
	router := SetupRouter(cfg)

	fmt.Println(">> Connecting to database ...")
	database.Init(cfg)

	fmt.Println(">> Setting up database ...")
	database.BuildDatabase()

	fmt.Println(">> Starting service ...")
	router.Run(cfg.Server.Port)

	//defer config.Close()
}
