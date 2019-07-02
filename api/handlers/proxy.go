package handlers

import (
	"encoding/json"
	"fmt"

	"github.com/kataras/iris"
	middleware "github.com/kimberly.luna/proxy-app/api/middleware"
)

// HandlerRedirection should redirect traffic
func HandlerRedirection(app *iris.Application) {
	app.Get("/ping", middleware.ProxyMiddleware, proxyHandler)
}

func proxyHandler(c iris.Context) {
	response, err := json.Marshal(middleware.Queue)

	fmt.Println("Result in proxy handler ", middleware.Queue)
	if err != nil {
		c.JSON(iris.Map{"status": 400, "result": "parse error"})
		return
	}
	c.JSON(iris.Map{"result": string(response)})
}
