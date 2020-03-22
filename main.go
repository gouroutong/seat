package main

import (
	"github.com/kataras/iris"
	"xProcessBackend/conf"
	"xProcessBackend/router"
)


func main() {
	conf.Init()
	configs := conf.GetConfig()
	app := router.Router()
	app.Run(iris.Addr(configs.Server), iris.WithCharset("UTF-8"))
}
