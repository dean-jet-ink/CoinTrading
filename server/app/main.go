package main

import (
	"cointrading/app/config"
	"cointrading/app/di"
	"fmt"
)

func init() {
	config.LoggingSettings(config.LogFileName())
}

func main() {
	// setupRouterというメソッドを作るべき
	// factoryでrouterをセットアップは変すぎる
	fmt.Println("Starting")
	router := di.Initialize()
	router.Run(fmt.Sprintf(":%v", config.Port()))
}
