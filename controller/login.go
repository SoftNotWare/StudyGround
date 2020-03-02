package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login 处理登录请求
func Login(ctx *gin.Context) {
	var param struct {
		Username string `json:"user"`
		Password string `json:"pwd"`
	}
	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}
	_ = param

}
