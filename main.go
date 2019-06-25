package main

import (
	handlers "github.com/kimberly.luna/proxy-app/api/handlers"
	middleware "github.com/kimberly.luna/proxy-app/api/middleware"
	server "github.com/kimberly.luna/proxy-app/api/server"
	utils "github.com/kimberly.luna/proxy-app/api/utils"
)

func main() {
	utils.LoadEnv()
	app := server.SetUp()
	middleware.InitQueue()
	handlers.HandlerRedirection(app)
	server.RunServer(app)
}
