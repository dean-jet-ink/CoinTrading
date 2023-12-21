package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load("/src/.env"); err != nil {
		log.Fatalf("Faild to load env file: %v", err)
	}
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
