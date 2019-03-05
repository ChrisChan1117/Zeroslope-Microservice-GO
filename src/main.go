package main

import (
	"fmt"

	"./database"
	"./router"
	"./utilities"
)

// @title ZeroSlope API
// @version 1.0
// @description Zeroslope microservice architecture written in go
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	fmt.Println(">> Loading configuration ...")
	cfg := utilities.LoadConfigFromPath("config.yaml")

	fmt.Println(">> Loading routes ...")
	router := router.SetupRouter(cfg)

	fmt.Println(">> Connecting to database ...")
	database.Init(cfg)

	fmt.Println(">> Setting up database ...")
	database.BuildDatabase()

	fmt.Println(">> Starting service ...")
	router.Run(cfg.Server.Port)

	//defer utilities.Close()
}
