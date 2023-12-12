package main

import (
	"cointrading/app/config"
	"log"
)

func main() {
	config.LoggingSettings(config.LogFileName())

	log.Println(config.BitflyerApiKey())
	log.Println(config.BitflyerApiSecret())
}
