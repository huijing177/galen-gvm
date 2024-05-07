package middleware

import (
	"galen-gvm/internal/jwtutils"
	"galen-gvm/utils"

	"github.com/gin-gonic/gin"
)

func JwtAuth(ctx *gin.Context) {
	token := utils.GetToken(ctx)
	if token == "" {
		utils.HTTPNoAuth("未登录或非法访问", ctx)
		ctx.Abort()
		return
	}
	jwt := jwtutils.NewJwt()
	claims, err := jwt.ParseToken(token)
	if err != nil {
		utils.HTTPNoAuth(err.Error(), ctx)
		utils.ClearToken(ctx)
		ctx.Abort()
		return
	}
	ctx.Set("claims", claims)
	ctx.Next()
}
