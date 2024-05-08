package request

type RegisterReq struct {
	Username  string `json:"userName" example:"用户名" binding:"required"`
	Password  string `json:"passWord" example:"密码" binding:"required"`
	NickName  string `json:"nickName" example:"昵称" binding:"required"`
	HeaderImg string `json:"headerImg" example:"头像链接"`
	Phone     string `json:"phone" example:"电话号码" binding:"required"`
	Email     string `json:"email" example:"电子邮箱" binding:"required"`
}
