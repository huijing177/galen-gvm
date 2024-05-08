package system

import (
	"galen-gvm/global"
	"galen-gvm/internal/jwtutils"
	"galen-gvm/model/request"
	"galen-gvm/model/system"
	"galen-gvm/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func UserRouter(c *gin.RouterGroup) {
	// 改密码只会维护一个版本
	c.POST("/changepassword", ChangePassword)
	// V1Group := c.Group("v1")
}

// ChangePassword
// @Tags     user
// @Summary  更改密码
// @Produce   application/json
// @Param    data  body      request.ChangePasswordReq                                             true  "原密码和新密码"
// @Success  200   {object}  utils.Response{msg=string}
// @Router   /system/user/changepassword [post]
func ChangePassword(ctx *gin.Context) {
	var req request.ChangePasswordReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		global.GVA_LOG.Error("ChangePasswordReq bind err:", zap.Error(err))
		utils.HTTPFail(nil, err.Error(), ctx)
		return
	}
	// TODO 密码格式验证

	id := jwtutils.GetUserIdByCtx(ctx)
	if id == 0 {
		utils.HTTPNoAuth("请登录", ctx)
		return
	}
	_, err := system.ChangePassword(&system.SysUser{GVA_MODEL: global.GVA_MODEL{ID: id}, Password: req.Password}, req.NewPassword)
	if err != nil {
		global.GVA_LOG.Error("system.ChangePassword err:", zap.Error(err))
		utils.HTTPFail(nil, "更改密码失败", ctx)
		return
	}
	utils.HTTPOk(nil, "修改成功", ctx)
}
