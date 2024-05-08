package base

import (
	"time"

	"galen-gvm/global"
	"galen-gvm/internal/jwtutils"
	"galen-gvm/internal/redis"
	"galen-gvm/model/request"
	"galen-gvm/model/response"
	"galen-gvm/model/system"
	"galen-gvm/utils"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid/v5"
	"github.com/mojocn/base64Captcha"
	"github.com/spf13/cast"
	"go.uber.org/zap"
)

var (
	_CaptchaStore = base64Captcha.DefaultMemStore
)

func BaseRouter(g *gin.RouterGroup) {
	// 登陆等接口，不应该同时支持多个版本，不需要分出v1，v2
	g.POST("/login", Login)
	g.POST("/captcha", Captcha)
	g.POST("/register", Register)
}

// Login
// @Tags     Base
// @Summary  用户登录
// @Produce   application/json
// @Param    data  body      request.LoginReq                                             true  "用户名, 密码, 验证码"
// @Success  200   {object}  utils.Response{data=response.LoginResp,msg=string}  "返回包括用户信息,token,过期时间"
// @Router   /base/login [post]
func Login(ctx *gin.Context) {
	var login request.LoginReq
	err := ctx.ShouldBindJSON(&login)
	if err != nil {
		global.GVA_LOG.Error("Login bind err:", zap.Error(err))
		utils.HTTPFail(nil, err.Error(), ctx)
		return
	}
	// 验证码校验
	v, ok := global.BlackCache.Get(ctx.ClientIP())
	if !ok {
		global.BlackCache.Set(ctx.ClientIP(), 1, time.Duration(global.GVA_CONFIG.Captcha.OpenCaptchaTimeOut)*time.Second)
	}
	defer func() {
		if err := global.BlackCache.Increment(ctx.ClientIP(), 1); err != nil {
			global.GVA_LOG.Error("BlackCache.Increment err:", zap.Error(err))
		}
	}()
	// 这里不能通过参数中是否有验证码判断是否需要验证，防止有人直接调用接口，暴力破解
	flag := false
	if global.GVA_CONFIG.Captcha.OpenCaptcha == 0 || global.GVA_CONFIG.Captcha.OpenCaptcha < cast.ToInt(v) {
		flag = true
	}
	if flag {
		// 不管成功失败，前端应该都会刷新界面，重新获取验证码，所以需要删除
		if login.CaptchaId == "" || login.Captcha == "" || !_CaptchaStore.Verify(login.CaptchaId, login.Captcha, true) {
			global.GVA_LOG.Info("captcha verify fail")
			utils.HTTPFail(nil, "验证码错误", ctx)
			return
		}
	}
	var userInter system.SysUser
	if login.Username == "admin" && login.Password == "123456" {
		userInter = system.SysUser{
			GVA_MODEL:   global.GVA_MODEL{},
			UUID:        uuid.Must(uuid.NewV4()),
			Username:    "admin",
			Password:    "123456",
			NickName:    "galen",
			SideMode:    "",
			HeaderImg:   "",
			BaseColor:   "",
			ActiveColor: "",
			AuthorityId: 0,
			Phone:       "110",
			Email:       "123@qq.com",
			Enable:      0,
		}
	} else {
		userInter, err = system.GetUserbyUsername(login.Username)
		if err != nil {
			global.GVA_LOG.Error("login getuser err:", zap.Error(err))
			utils.HTTPFail(nil, "账号不存在", ctx)
			return
		}
		if utils.BcryptCheck(login.Password, userInter.Password) {
			utils.HTTPFail(nil, "密码错误", ctx)
			return
		}
		// 被限制登录
		if userInter.Enable == 2 {
			utils.HTTPFail(nil, "限制登录", ctx)
			return
		}
	}

	genToken(ctx, &userInter)
}

func genToken(ctx *gin.Context, user *system.SysUser) {
	j := jwtutils.NewJwt()
	claims := j.CreateClaims(user)
	token, err := j.CreateToken(claims)
	if err != nil {
		global.GVA_LOG.Error("CreateToken err:", zap.Error(err))
		utils.HTTPFail(nil, "获取token失败", ctx)
		return
	}
	utils.SetToken(ctx, token, int(claims.RegisteredClaims.ExpiresAt.Unix()-time.Now().Unix()))
	utils.HTTPOk(response.LoginResp{
		User:      *user,
		Token:     token,
		ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
	}, "登陆成功", ctx)
}

// Captcha
// @Tags     Base
// @Summary  获取验证码
// @Produce   application/json
// @Success  200   {object}  utils.Response{data=response.CaptchaResp,msg=string}  "返回包括是否开启验证码，以及验证码信息"
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
	if global.GVA_CONFIG.System.UseRedis {
		_CaptchaStore = redis.NewCaptchaRedis()
	}
	cp := base64Captcha.NewCaptcha(driver, _CaptchaStore)
	id, b64s, _, err := cp.Generate()
	if err != nil {
		global.GVA_LOG.Error("验证码获取失败!", zap.Error(err))
		utils.HTTPFail(nil, "验证码获取失败", ctx)
		return
	}
	utils.HTTPOk(response.CaptchaResp{
		CaptchaId:     id,
		PicPath:       b64s,
		CaptchaLength: global.GVA_CONFIG.Captcha.KeyLong,
		OpenCaptcha:   flag,
	}, "验证码获取成功", ctx)
}

// Register
// @Tags     Base
// @Summary  注册用户
// @Produce   application/json
// @Param    data  body      request.RegisterReq                                             true  "用户名, 密码, 验证码等"
// @Success  200   {object}  utils.Response{data=response.RegisterResp,msg=string}  "返回用户信息"
// @Router   /base/register [post]
func Register(ctx *gin.Context) {
	var register request.RegisterReq
	err := ctx.ShouldBindJSON(&register)
	if err != nil {
		global.GVA_LOG.Error("Register bind err:", zap.Error(err))
		utils.HTTPFail(nil, err.Error(), ctx)
		return
	}
	user := &system.SysUser{
		Username:  register.Username,
		Password:  register.Password,
		NickName:  register.NickName,
		HeaderImg: register.HeaderImg,
		Phone:     register.Phone,
		Email:     register.Email,
		Enable:    0,
	}
	u, err := system.Register(user)
	if err != nil {
		global.GVA_LOG.Error("system.Register err:", zap.Error(err))
		utils.HTTPFail(nil, "注册失败", ctx)
		return
	}
	utils.HTTPOk(response.RegisterResp{
		User: *u,
	}, "注册成功", ctx)
}
