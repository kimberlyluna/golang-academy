package main

import (
	handlers "github.com/kimberly.luna/proxy-app/api/handlers"
	server "github.com/kimberly.luna/proxy-app/api/server"
	utils "github.com/kimberly.luna/proxy-app/api/utils"
)

func main() {
	utils.LoadEnv()
	app := server.SetUp()
	handlers.HandlerRedirection(app)
	server.RunServer(app)
}
