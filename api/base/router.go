package base

import (
	"galen-gvm/global"
	"galen-gvm/model/request"
	"galen-gvm/model/system"
	"galen-gvm/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func BaseRouter(g *gin.RouterGroup) {
	// 登陆等接口，不应该同时支持多个版本，不需要分出v1，v2
	g.POST("/login", Login)

}

func Login(ctx *gin.Context) {
	var login request.Login
	err := ctx.ShouldBindJSON(&login)
	if err != nil {
		global.GVA_LOG.Error("Login bind err:", zap.Error(err))
		utils.HTTPFail(nil, err.Error(), ctx)
		return
	}
	// TODO 验证码校验
	userInter, err := system.GetUserbyUsername(login.Username)
	if err != nil {
		global.GVA_LOG.Error("login getuser err:", zap.Error(err))
		utils.HTTPFail(nil, "账号不存在", ctx)
		return
	}
	if login.Password != userInter.Password {
		utils.HTTPFail(nil, "密码错误", ctx)
		return
	}
	// 被限制登录
	if userInter.Enable == 2 {
		utils.HTTPFail(nil, "限制登录", ctx)
		return
	}

	genToken(ctx, &userInter)
}

func genToken(ctx *gin.Context, user *system.SysUser) {
	// TODO
}
