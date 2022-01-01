package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func BindPayload(ctx *gin.Context, obj interface{}) error {
	objAddr := &obj
	bound := binding.Default(ctx.Request.Method, ctx.ContentType())
	return ctx.ShouldBindWith(objAddr, bound)
}
