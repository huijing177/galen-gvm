package request

type LoginReq struct {
	Username  string `json:"username" binding:"required"` // 用户名
	Password  string `json:"password" binding:"required"` // 密码
	Captcha   string `json:"captcha"`                     // 验证码
	CaptchaId string `json:"captchaId"`                   // 验证码ID
}
