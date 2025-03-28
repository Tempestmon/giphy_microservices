package main

import (
	apiservice "giphy_microservices/api_service"
)

func main() {
	// app := app.InitApp()
	// if app.HttpServer == nil {
	// 	log.Fatal("Server is nil - Wire injection failed")
	// }
	// log.Fatal(app.HttpServer.Start())
	apiservice.NewServer("127.0.0.1:8080").Start()
}
