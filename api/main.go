package main

import (
	"api/config"
	"api/view"
)

func main() {
	view.NewServer(
		config.ServerConfig{
			IP:   "127.0.0.1",
			Port: "8080",
		},
	).Start()
}
