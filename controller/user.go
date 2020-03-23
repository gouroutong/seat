package controller

import (
	"context"
	"xProcessBackend/model"
	"xProcessBackend/serializer"
)

func Login(ctx context.Context){
	var user model.User
	ctx.ReadJSON(&user)
	ctx.JSON(serializer.GetResponse(&user))
}