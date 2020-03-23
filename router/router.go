package router

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/core/router"
	"xProcessBackend/controller"

	//"github.com/kataras/iris/core/router"
)

func Router() *iris.Application {
	app := iris.Default()

	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"PUT", "PATCH", "GET", "POST", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
		AllowCredentials: true,
	})

	router1 := app.Party("/api", crs).
		AllowMethods(iris.MethodOptions)
	{
		router1.PartyFunc("/seat", func(process router.Party) {
			process.Get("/test", func(c context.Context) {
				c.JSON(iris.Map{"name":"zhangsan"})
			})
		})

		router1.PartyFunc("/user", func(user router.Party) {
			user.Post("/Login",controller.Login)

		})
	}
	return app
}
