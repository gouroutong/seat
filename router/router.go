package router

import (
  "github.com/iris-contrib/middleware/cors"
  "github.com/kataras/iris"
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
    router1.PartyFunc("/room", func(room router.Party) {
      room.Post("/get", controller.GetRoom)
      room.Post("/new", controller.NewRoom)
    })

    router1.PartyFunc("/user", func(user router.Party) {
      user.Post("/login", controller.Login)
      user.Post("/new", controller.New)
    })

    router1.PartyFunc("/order", func(order router.Party) {
      order.Post("/get_all", controller.OrderList)
      order.Post("/new", controller.NewOrder)
    })
  }
  return app
}
