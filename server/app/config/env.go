package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func init() {
	path := fmt.Sprintf("%s/.env", DockerWorkDir())

	if err := godotenv.Load(path); err != nil {
		log.Fatalf("Faild to load env file: %v", err)
	}
}

func FrameWork() int {
	frameworkStr := os.Getenv("FRAMEWORK")

	framework, err := strconv.Atoi(frameworkStr)
	if err != nil {
		log.Println("Invalid or missing FRAMEWORK. Defaulting to 1")
		return 1
	}

	return framework
}

func DB() int {
	dbStr := os.Getenv("DB")

	db, err := strconv.Atoi(dbStr)
	if err != nil {
		log.Println("Invalid or missing DB. Defaulting to 1")
		return 1
	}

	return db
}

func ORM() int {
	ormStr := os.Getenv("ORM")

	orm, err := strconv.Atoi(ormStr)
	if err != nil {
		log.Println("Invalid or missing ORM. Defaulting to 1")
		return 1
	}

	return orm
}

func FrontEndURL() string {
	return os.Getenv("FRONT_END_URL")
}

func Port() string {
	return os.Getenv("PORT")
}

func LogFileName() string {
	return os.Getenv("LOG_FILE_NAME")
}

func DockerWorkDir() string {
	return os.Getenv("DOCKER_WORKDIR")
}

func GetCandleLimit() int {
	limitStr := os.Getenv("GET_CANDLE_LIMIT")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		log.Println("Invalid or missing GET_CANDLE_LIMIT. Defaulting to 50")
		return 50
	}

	return limit
}

func BitflyerApiKey() string {
	return os.Getenv("BITFLYER_API_KEY")
}

func BitflyerApiSecret() string {
	return os.Getenv("BITFLYER_API_SECRET")
}

func WebSocketHandshakeTimeout() int {
	timeoutStr := os.Getenv("WEB_SOCKET_HANDSHAKE_TIMEOUT")

	timeout, err := strconv.Atoi(timeoutStr)
	if err != nil {
		log.Println("Invalid or missing WEB_SOCKET_HANDSHAKE_TIMEOUT. Defaulting to 45")
		return 45
	}

	return timeout
}

func WebSocketBufferSize() int {
	bufferStr := os.Getenv("WEB_SOCKET_BUFFER_SIZE")

	buffer, err := strconv.Atoi(bufferStr)
	if err != nil {
		log.Println("Invalid or missing WEB_SOCKET_BUFFER_SIZE. Defaulting to 1024")
		return 1024
	}

	return buffer
}
