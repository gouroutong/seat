package controller

import (
  "github.com/kataras/iris/context"
  "xProcessBackend/model"
  "xProcessBackend/serializer"
)

func Login(ctx context.Context) {
  var user model.User
  var user1 model.User
  ctx.ReadJSON(&user)
  err := user.Login(&user1)
  ctx.JSON(serializer.GetResponse(user1, err))
}
func New(ctx context.Context) {
  var (
    user model.User
  )
  ctx.ReadJSON(&user)
  ctx.JSON(serializer.GetResponse(user, user.New()))
}
