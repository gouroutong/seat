package controller

import (
  "github.com/kataras/iris"
  "xProcessBackend/model"
  "xProcessBackend/serializer"
)

func NewOrder(ctx iris.Context) {
  var p model.Order
  ctx.ReadJSON(&p)
  err := p.NewOrder()
  ctx.JSON(serializer.GetResponse(p, err))
}

func OrderList(ctx iris.Context) {
  var (
    p struct {
      SeatId int64 `json:"seat_id"`
    }
    list []model.Order
  )
  ctx.ReadJSON(&p)
  err := model.OrderList(p.SeatId, &list)
  ctx.JSON(serializer.GetResponse(list, err))
}
