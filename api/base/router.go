package base

import (
	"time"

	"galen-gvm/global"
	"galen-gvm/internal/redis"
	"galen-gvm/model/request"
	"galen-gvm/model/response"
	"galen-gvm/model/system"
	"galen-gvm/utils"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"github.com/spf13/cast"
	"go.uber.org/zap"
)

func BaseRouter(g *gin.RouterGroup) {
	// 登陆等接口，不应该同时支持多个版本，不需要分出v1，v2
	g.POST("/login", Login)
	g.POST("/captcha", Captcha)

}

// Login
// @Tags     Base
// @Summary  用户登录
// @Produce   application/json
// @Param    data  body      request.Login                                             true  "用户名, 密码, 验证码"
// @Success  200   {object}  utils.Response{data=response.LoginResponse,msg=string}  "返回包括用户信息,token,过期时间"
// @Router   /base/login [post]
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

// Captcha
// @Tags     Base
// @Summary  获取验证码
// @Produce   application/json
// @Success  200   {object}  utils.Response{data=response.CaptchaResponse,msg=string}  "返回包括是否开启验证码，以及验证码信息"
// @Router   /base/captcha [post]
func Captcha(ctx *gin.Context) {
	// 不管是否开启，都需要回传给前端，并把是否开启也回传回去，让前端抉择
	open := global.GVA_CONFIG.Captcha.OpenCaptcha
	openTimeOut := global.GVA_CONFIG.Captcha.OpenCaptchaTimeOut
	cliIP := ctx.ClientIP()

	v, ok := global.BlackCache.Get(cliIP)
	if !ok {
		global.BlackCache.Set(cliIP, 1, time.Second*time.Duration(openTimeOut))
	}
	flag := false
	if open == 0 || open < cast.ToInt(v) {
		flag = true
	}

	driver := base64Captcha.NewDriverDigit(global.GVA_CONFIG.Captcha.ImgHeight, global.GVA_CONFIG.Captcha.ImgWidth, global.GVA_CONFIG.Captcha.KeyLong, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, base64Captcha.DefaultMemStore)
	if global.GVA_CONFIG.System.UseRedis {
		store := redis.NewCaptchaRedis()
		cp = base64Captcha.NewCaptcha(driver, store)
	}
	id, b64s, _, err := cp.Generate()
	if err != nil {
		global.GVA_LOG.Error("验证码获取失败!", zap.Error(err))
		utils.HTTPFail(nil, "验证码获取失败", ctx)
		return
	}
	utils.HTTPOk(response.CaptchaResponse{
		CaptchaId:     id,
		PicPath:       b64s,
		CaptchaLength: global.GVA_CONFIG.Captcha.KeyLong,
		OpenCaptcha:   flag,
	}, "验证码获取成功", ctx)
}
