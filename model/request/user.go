package request

type ChangePasswordReq struct {
	Password    string `json:"password" binding:"required"`    // 原密码
	NewPassword string `json:"newpassword" binding:"required"` // 新密码
}
