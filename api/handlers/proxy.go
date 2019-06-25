package handlers

import "github.com/kataras/iris"
import middleware "github.com/kimberly.luna/proxy-app/api/middleware"

// HandlerRedirection should redirect traffic
func HandlerRedirection(app *iris.Application) {
	app.Get("/ping", middleware.ProxyMiddleware, proxyHandler)
}

func proxyHandler(c iris.Context) {
	c.JSON(iris.Map{"result": "ok"})
}
