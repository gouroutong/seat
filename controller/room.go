package controller

import (
  "github.com/kataras/iris"
  "xProcessBackend/model"
  "xProcessBackend/serializer"
)

func GetRoom(ctx iris.Context) {
  var p model.Room
  ctx.ReadJSON(&p)
  err := p.GetRoom()
  ctx.JSON(serializer.GetResponse(p, err))
}

func NewRoom(ctx iris.Context)  {
  var p model.Room
  ctx.ReadJSON(&p)
  err := p.NewRoom()
  ctx.JSON(serializer.GetResponse(p, err))
}
