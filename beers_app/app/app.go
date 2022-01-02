package app

import (
	"os"
	"strconv"
)

func Run() error {
	serverHost := os.Getenv("BACKEND_HOST")
	serverPort, err := strconv.ParseUint(os.Getenv("BACKEND_PORT"), 10, 0)
	if err != nil {
		panic("Environment variable BACKEND_PORT is not defined")
	}
	server := NewServer(serverHost, uint(serverPort))
	return server.Run()
}
