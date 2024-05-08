package response

type CaptchaResp struct {
	CaptchaId     string `json:"captchaId"`
	PicPath       string `json:"picPath"`
	CaptchaLength int    `json:"captchaLength"`
	OpenCaptcha   bool   `json:"openCaptcha"`
}
